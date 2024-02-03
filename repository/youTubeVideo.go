package repository

import (
	"FamPayApp/contracts"
	"FamPayApp/model"
	"database/sql"
	"github.com/google/uuid"
	"log"
	"time"
)

type YoutubeVideoRepository struct {
	DB *sql.DB
}

func NewYoutubeVideoRepository(db *sql.DB) YoutubeVideoRepositoryInterface {
	return &YoutubeVideoRepository{DB: db}
}

// InsertYoutubeVideo implements YoutubeVideoRepositoryInterface
func (m *YoutubeVideoRepository) InsertYoutubeVideo(videos []model.YoutubeVideo) error {
	for _, video := range videos{
		_, err := m.DB.Exec("INSERT INTO youtube_video (title, category, description, video_id, video_link, published_at, thumbnail) VALUES ($1, $2, $3, $4, $5, $6, $7)", video.Title, video.Category, video.Description, video.VideoId, video.VideoLink, video.PublishedAt, video.Thumbnail)
		if err != nil {
			log.Println("error while insertion: ", err)
			return err
		}
	}
	return nil
}
// GetAllYoutubeVideos implements YoutubeVideoRepositoryInterface
func (m *YoutubeVideoRepository) GetAllYoutubeVideos() (videos []contracts.GetVideoResponse, err error) {
	row, err := m.DB.Query("SELECT title, description, video_link, published_at, thumbnail FROM youtube_video ORDER BY published_at DESC")
	if err != nil {
		log.Println(err)
		return
	}

	for row.Next() {
		var (
			title        string
			description  string
			video_link   string
			published_at time.Time
			thumbnail    string
		)
		err = row.Scan(&title, &description, &video_link, &published_at, &thumbnail)
		if err != nil {
			log.Println(err)
		} else {
			video := contracts.GetVideoResponse{Title: title, Description: description, VideoLink: video_link, PublishedAt: published_at, Thumbnail: thumbnail}
			videos = append(videos, video)
		}
	}
	return
}

// SearchYoutubeVideo implements YoutubeVideoRepositoryInterface
func (m *YoutubeVideoRepository) SearchYoutubeVideo(title, description string) (video contracts.GetVideoResponse) {
	row, err := m.DB.Query("SELECT title, description, video_link, published_at, thumbnail FROM youtube_video WHERE title = $1 and description = $2", title, description)
	if err != nil {
		log.Println(err)
		return
	}

	for row.Next() {
		var (
			title        string
			description  string
			video_link   string
			published_at time.Time
			thumbnail    string
		)
		err := row.Scan(&title, &description, &video_link, &published_at, &thumbnail)
		if err != nil {
			log.Println(err)
		} else {
			video = contracts.GetVideoResponse{Title: title, Description: description, VideoLink: video_link, PublishedAt: published_at, Thumbnail: thumbnail}
		}
	}
	return
}

// GetYoutubeVideoByVideoId implements YoutubeVideoRepositoryInterface
func (m *YoutubeVideoRepository) GetYoutubeVideoByVideoId(videoId string) (video model.YoutubeVideo) {
	row, err := m.DB.Query("SELECT * FROM youtube_video WHERE video_id = $1", videoId)
	if err != nil {
		log.Println(err)
		return
	}

	for row.Next() {
		var (
			id           uuid.UUID
			title        string
			category     string
			description  string
			video_id     string
			video_link   string
			published_at time.Time
			thumbnail    string
		)
		err := row.Scan(&id, &title, &category, &description, &video_id, &video_link, &published_at, &thumbnail)
		if err != nil {
			log.Println(err)
		} else {
			video = model.YoutubeVideo{ID: id, Title: title, Category: category, Description: description, VideoId: video_id, VideoLink: video_link, PublishedAt: published_at, Thumbnail: thumbnail}
		}
	}
	return
}

