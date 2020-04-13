package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
	"github.com/silverswords/fake/app/date/model"
)

//Month is to run month
func Month(router *gin.Engine) {
	router.GET("/month", month)
}

func month(c *gin.Context) {
	month := GetMonth()

	c.String(200, "month: %s", month)
}

//GetMonth is to get month
func GetMonth() string {
	month := rand.GetRand().StringArray(&model.MonthData)

	return month
}
