package Auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vi350/spbstu-hackathon-autumn19/DB"
	"github.com/vi350/spbstu-hackathon-autumn19/Model"
	"log"
)

func Auth(c *gin.Context) {

	var status int
	var message string

	var dataInUser Model.DataInUser
	var user Model.User
	var token string
	var skills []string
	var favourites []int
	var ignored []int

	err := c.ShouldBindJSON(&dataInUser)
	if err != nil {
		status = 400
		message = "maybe not all parameters are specified: " + err.Error()
		log.Fatalln(err.Error())
	} else {
		status = 200
		message = "seems to be ok"
		err = DB.DB.Model(&user).Where("uniqueid = ?", dataInUser.Id).Select()
		user.Name = dataInUser.FirstName
		user.Uniqueid = dataInUser.Id
		user.Username = dataInUser.Username
		user.Photourl = dataInUser.PhotoUrl
		user.Rating = 5
		user.Busy = false
		if err != nil {
			err = DB.DB.Insert(&user)
		}
		token, _ = GenerateRandomString(15)
		_, _ = DB.DB.Model(&user).Set("token = ?", token).Where("uniqueid = ?", dataInUser.Id).Update()
		_ = json.Unmarshal([]byte(user.Skills),&skills)
		_ = json.Unmarshal([]byte(user.Favourites),&favourites)
		_ = json.Unmarshal([]byte(user.Ignored),&ignored)
	}

	c.JSON(status, gin.H{
		"message": message,
		"token": token,
		"uniqueid": user.Uniqueid,
		"name": user.Name,
		"photourl": user.Photourl,
		"skills": skills,
		"favourites": favourites,
		"ignored": ignored,
		"busy": user.Busy,
	})
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}