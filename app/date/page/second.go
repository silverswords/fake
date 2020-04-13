package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
)

//Second is to run second
func Second(router *gin.Engine) {
	router.GET("/second", second)
}

func second(c *gin.Context) {
	second := GetSecond()

	c.String(200, "second: %d", second)
}

//GetSecond is to get second
func GetSecond() int {
	second := rand.GetRand().RandRange(0, 60)

	return second
}
