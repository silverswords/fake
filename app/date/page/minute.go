package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
	"github.com/silverswords/fake/app/date/model"
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
	var minute int

	if model.IsConfig("minute") {
		minMinute, maxMinute := getMinuteData()
		minute = rand.GetRand().RandRange(minMinute, maxMinute)
	} else {
		minute = rand.GetRand().RandRange(0, 59)
	}

	return minute
}

func getMinuteData() (minMinute int, maxMinute int) {
	dateMap := model.GetDateConfig()
	minuteMap := dateMap["minute"].(map[string]interface{})
	minMinute = int(minuteMap["minminute"].(float64))
	maxMinute = int(minuteMap["maxminute"].(float64))
	return
}
