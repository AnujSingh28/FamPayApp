package controller

import "github.com/gin-gonic/gin"

type YoutubeVideoControllerInterface interface {
	GetYoutubeVideo(g *gin.Context)
	FetchYoutubeVideos(g *gin.Context)
	GetAllYoutubeVideos(g *gin.Context)
}
