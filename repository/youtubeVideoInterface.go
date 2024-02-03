package repository

import (
	"FamPayApp/contracts"
	"FamPayApp/model"
)

type YoutubeVideoRepositoryInterface interface {
	SearchYoutubeVideo(title, description string) contracts.GetVideoResponse
	InsertYoutubeVideo(post []model.YoutubeVideo) error
	GetAllYoutubeVideos() ([]contracts.GetVideoResponse, error)
	GetYoutubeVideoByVideoId(videoId string) (video model.YoutubeVideo)
}
