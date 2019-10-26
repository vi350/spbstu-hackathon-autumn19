package Update

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Busy(c *gin.Context) {
	message := "ok"
	status := 200

	_, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
	} else {
		c.JSON(status, gin.H{
			"message": message,
		})
	}
}
