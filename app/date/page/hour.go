package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
)

//Hour is to run hour
func Hour(router *gin.Engine) {
	router.GET("/hour", hour)
}

func hour(c *gin.Context) {
	hour := GetHour()

	c.String(200, "hour: %d", hour)
}

//GetHour is to get hour
func GetHour() int {
	hour := rand.GetRand().RandRange(0, 24)

	return hour
}
