package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
)

//Minute is to run minute
func Minute(router *gin.Engine) {
	router.GET("/minute", minute)
}

func minute(c *gin.Context) {
	minute := GetMinute()

	c.String(200, "minute: %d", minute)
}

//GetMinute is to get minute
func GetMinute() int {
	minute := rand.GetRand().RandRange(0, 60)

	return minute
}
