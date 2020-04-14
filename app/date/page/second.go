package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
	"github.com/silverswords/fake/app/date/model"
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
	var second int

	if model.IsConfig("second") {
		minSecond, maxSecond := getSecondData()
		second = rand.GetRand().RandRange(minSecond, maxSecond)
	} else {
		second = rand.GetRand().RandRange(0, 59)
	}

	return second
}

func getSecondData() (minSecond int, maxSecond int) {
	dateMap := model.GetDateConfig()
	secondMap := dateMap["second"].(map[string]interface{})
	minSecond = int(secondMap["minsecond"].(float64))
	maxSecond = int(secondMap["maxsecond"].(float64))
	return
}
