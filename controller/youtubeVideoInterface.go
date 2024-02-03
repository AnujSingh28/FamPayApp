package controller

import "github.com/gin-gonic/gin"

type YoutubeVideoControllerInterface interface {
	//InsertYoutubeVideo(g *gin.Context)
	GetYoutubeVideo(g *gin.Context)
	FetchYoutubeVideos(g *gin.Context)
	GetAllYoutubeVideos(g *gin.Context)
}
