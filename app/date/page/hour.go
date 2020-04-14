package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
	"github.com/silverswords/fake/app/date/model"
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
	var hour int

	if model.IsConfig("hour") {
		minHour, maxHour := getHourData()
		hour = rand.GetRand().RandRange(minHour, maxHour)
	} else {
		hour = rand.GetRand().RandRange(0, 23)
	}

	return hour
}

func getHourData() (minHour int, maxHour int) {
	dateMap := model.GetDateConfig()
	hourMap := dateMap["hour"].(map[string]interface{})
	minHour = int(hourMap["minhour"].(float64))
	maxHour = int(hourMap["maxhour"].(float64))
	return
}
