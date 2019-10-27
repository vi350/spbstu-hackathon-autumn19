package Update

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vi350/spbstu-hackathon-autumn19/DB"
	"github.com/vi350/spbstu-hackathon-autumn19/Model"
	"log"
)

func Busy(c *gin.Context) {
	fmt.Println("1")
	message := "ok"
	status := 200
	type Updatedata struct {
		Token   string  `json:"token"  binding:"required"`
		Busy    bool    `json:"busy"  binding:"required"`
	}
	var data Updatedata
	err := c.ShouldBindJSON(&data)
	var user Model.User
	fmt.Println("2")
	err = DB.DB.Model(&user).Where("token = ?", data.Token).Select()
	fmt.Println("3")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("4")
		_, _ = DB.DB.Model(&user).Set("busy = ?", data.Busy).Where("token = ?", data.Token).Update()
		fmt.Println("5")
		c.JSON(status, gin.H{
			"message": message,
			"busy": data.Busy,
		})
	}
}
