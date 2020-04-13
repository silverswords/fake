package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
)

//Day is to run day
func Day(router *gin.Engine) {
	router.GET("/day", day)
}

func day(c *gin.Context) {
	day := GetDay(30)

	c.String(200, "day: %d", day)
}

//GetDay is to get day with max day
func GetDay(maxDay int) int {
	day := rand.GetRand().RandRange(1, maxDay)

	return day
}
