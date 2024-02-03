package app

import (
	"FamPayApp/controller"
	"database/sql"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Routes *gin.Engine
}

func (a *App) CreateConnection(){
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", Database, Password, HostDB, DBNAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
}

func (a *App) CreateRoutes() {
	routes := gin.Default()
	controller := controller.NewYoutubeVideoController(a.DB)
	routes.GET("/getVideo", controller.GetYoutubeVideo)
	routes.GET("/allVideos", controller.GetAllYoutubeVideos)
	a.Routes = routes
}

func (a *App) StartCronJob() *cron.Cron {
	// Run the cron job to fetch youtube data every minute
	c := cron.New()
	_, err := c.AddFunc("* * * * *", func() {
		controller := controller.NewYoutubeVideoController(a.DB)
		controller.FetchYoutubeVideos(nil)
	})
	if err != nil {
		log.Fatalf("Error setting up cron job: %v", err)
	}
	c.Start()
	return c
}

func (a *App) Run(){
	a.Routes.Run(":8080")
}

