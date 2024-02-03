package controller

import (
	"FamPayApp/model"
	"FamPayApp/repository"
	"FamPayApp/utils"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type YoutubeVideoController struct {
	DB *sql.DB
}

func NewYoutubeVideoController(db *sql.DB) YoutubeVideoControllerInterface {
	return &YoutubeVideoController{DB: db}
}

// GetYoutubeVideo implements YoutubeVideoControllerInterface
func (m *YoutubeVideoController) GetYoutubeVideo(g *gin.Context) {
	title := g.Query("title")
	description := g.Query("description")
	if title == "" {
		g.JSON(400, gin.H{"msg": "Title can not be null"})
	}
	db := m.DB
	youtubeRepo := repository.NewYoutubeVideoRepository(db)
	video := youtubeRepo.SearchYoutubeVideo(title, description)
	if video.Title == "" {
		log.Println("No video exists")
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "No such video exists"})
		return
	}
	g.JSON(200, gin.H{"status": "success", "data": video, "msg": "GetYoutubeVideo Successful"})
}

func (m *YoutubeVideoController) GetAllYoutubeVideos(g *gin.Context){
	page := g.Query("page")
	if page == "" {
		g.JSON(400, gin.H{"msg": "page can not be null"})
	}
	pageNumber, _ := strconv.ParseInt(page, 10, 64)
	_recordsPerPage := g.Query("recordsPerPage")
	if _recordsPerPage == "" {
		g.JSON(400, gin.H{"msg": "recordsPerPage can not be null"})
	}
	recordsPerPage, _ := strconv.ParseInt(_recordsPerPage, 10, 64)

	db := m.DB
	youtubeRepo := repository.NewYoutubeVideoRepository(db)
	videos, err := youtubeRepo.GetAllYoutubeVideos()
	if err != nil {
		g.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	if len(videos) == 0 {
		log.Println("No video exists")
		g.JSON(500, gin.H{"status": "success", "data": nil, "msg": "No video exists"})
		return
	}
	tmpvideos := make([]interface{}, 0)
	for _, video := range videos {
		tmpvideos = append(tmpvideos, video)
	}
	paginatedVideos := utils.Paginate(tmpvideos, pageNumber, recordsPerPage)
	g.JSON(200, gin.H{"status": "success", "data": paginatedVideos, "msg": "GetYoutubeVideo Successful"})
}

func (m *YoutubeVideoController) FetchYoutubeVideos(g *gin.Context) {

	db := m.DB
	fetchedYoutubeVideos := fetchYouTubeData()
	newYoutubeVideos := findNewVideosOutOfFetched(db, fetchedYoutubeVideos)
	youtubeRepo := repository.NewYoutubeVideoRepository(db)
	err := youtubeRepo.InsertYoutubeVideo(newYoutubeVideos)
	if err != nil {
		log.Println("error while creating creating db entry: ", err)
	}
}

func fetchYouTubeData() (videos []model.YoutubeVideo){
	service, err := youtube.NewService(context.Background(), option.WithAPIKey("AIzaSyBS2HIC7Cz3rtM7Iv8R6vhsefzWSMCAx0c"))
	if err != nil {
		log.Println("Failed to create YouTube client")
		return
	}

	category := "cricket"
	call := service.Search.List([]string{"snippet"}).Q(category).Type("video").MaxResults(10)
	response, err := call.Do()
	if err != nil {
		log.Println("Failed to fetch YouTube videos", err, "--------", response)
		return
	}

	// Extract relevant data from the API response
	for _, item := range response.Items {
		//published at time layout : "2006-01-02T15:04:05.000Z"
		parsedTime, errr := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		log.Println("original time: ", item.Snippet.PublishedAt, "------", parsedTime, "-----------", errr)
		video := model.YoutubeVideo{
			Category:       category,
			Title:          item.Snippet.Title,
			Description:    item.Snippet.Description,
			Thumbnail:      item.Snippet.Thumbnails.Default.Url,
			VideoId:        item.Id.VideoId,
			PublishedAt:    parsedTime,
			VideoLink:      fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId),
		}
		videos = append(videos, video)
	}
	return
}

func findNewVideosOutOfFetched(db *sql.DB, fetchedVideos []model.YoutubeVideo) (newVideos []model.YoutubeVideo) {
	youtubeRepo := repository.NewYoutubeVideoRepository(db)
	for _, fetchedVideo := range fetchedVideos {
		video := youtubeRepo.GetYoutubeVideoByVideoId(fetchedVideo.VideoId)
		if video.VideoLink == "" {
			newVideos = append(newVideos, fetchedVideo)
		}
	}
	return
}