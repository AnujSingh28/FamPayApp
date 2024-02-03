package repository

import (
	"FamPayApp/contracts"
	"FamPayApp/model"
)

type YoutubeVideoRepositoryInterface interface {
	SearchYoutubeVideo(slug string) []contracts.GetVideoResponse
	InsertYoutubeVideo(post []model.YoutubeVideo) error
	GetAllYoutubeVideos() ([]contracts.GetVideoResponse, error)
	GetYoutubeVideoByVideoId(videoId string) (video model.YoutubeVideo)
}
