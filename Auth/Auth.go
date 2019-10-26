package Auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vi350/spbstu-hackathon-autumn19/Model"
	"github.com/vi350/spbstu-hackathon-autumn19/DB"
	"log"
)

func Auth(c *gin.Context) {

	var status int
	var message string

	var dataInUser Model.DataInUser
	var user Model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		status = 400
		message = "maybe not all parameters are specified: " + err.Error()
		log.Fatalln(err.Error())
	} else {
		status = 200
		message = "seems to be ok"
		err = DB.DB.Model(&user).Where("uniqueid = ?", dataInUser.Id).Select()
		if err != nil {
			err = DB.DB.Insert(&user)
		}
	}

	c.JSON(status, gin.H{
		"message":       message,
	})
}