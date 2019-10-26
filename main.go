package main

import (
<<<<<<< HEAD
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-contrib/cors"
=======
>>>>>>> 3cce0e48ae888be3083202ef3842a30631ebf0f3
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

<<<<<<< HEAD
	//var user Model.User
	//var token string
	//token, _ = GenerateRandomString(15)
	//_, _ = DB.DB.Model(&user).Set("token = ?", token).Where("id = ?", "321").Returning("*").Update()
	//
	//err = app.Run(":8080")
	//if err != nil {
	//	log.Panic(err)
	//}
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
=======
	DB.SelectBySkills([]string{"go","gin"},"pdrs")


	//log.Fatal(autotls.Run(app, "example1.com", "example2.com"))

	err = app.Run(":8080")
>>>>>>> 3cce0e48ae888be3083202ef3842a30631ebf0f3
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
