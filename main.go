package main

import (
	"FamPayApp/app"
)

func main() {
	var a app.App
	a.CreateConnection()
	//a.Migrate()
	a.CreateRoutes()
	c := a.StartCronJob()
	defer c.Stop()
	a.Run()
}
