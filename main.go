package main

import (
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vi350/spbstu-hackathon-autumn19/Auth"
	"github.com/vi350/spbstu-hackathon-autumn19/Basics"
	"github.com/vi350/spbstu-hackathon-autumn19/DB"
	"github.com/vi350/spbstu-hackathon-autumn19/Update"
	"io"
	"log"
	"os"
)

func main() {

	// логирование (для heroku не нужен)
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Panic(err)
	}

	app := gin.Default()

	// Logging gin output in log.txt
	gin.DefaultWriter = io.MultiWriter(logFile)
	log.SetOutput(logFile)

	//app.Use(cors.Default())

	app.GET("/", Basics.Welcome)
	app.POST("/auth", Auth.Auth)
	app.GET("/busy/:data", Update.Busy)

	DB.ConnectDB()
	DB.CreateTables()

	DB.SelectBySkills([]string{"go","vue"},"pdrs")


	//log.Fatal(autotls.Run(app, "example1.com", "example2.com"))

	err = app.Run(":8080")
	if err != nil {
		log.Panic(err)
	}
}