package Auth

import (
	"crypto/rand"
	"encoding/base64"
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
		user.Rating = 5
		user.Busy = false
		if err != nil {
			err = DB.DB.Insert(&user)
		}
		var token string
		token, _ = GenerateRandomString(15)
		_, _ = DB.DB.Model(&user).Set("token = ?", token).Where("uniqueid = ?", dataInUser.Id).Update()
	}

	c.JSON(status, gin.H{
		"message": message,
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