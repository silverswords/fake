package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
	"github.com/silverswords/fake/app/date/model"
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
	var day int

	if getDayConfig() {
		minDay, maxDay := getDayData()
		day = rand.GetRand().RandRange(minDay, maxDay)
	} else {
		day = rand.GetRand().RandRange(1, maxDay)
	}

	return day
}

func getDayConfig() bool {
	dateMap := model.GetDateConfig()

	for k := range dateMap {
		if k == "day" {
			return true
		}
	}
	return false
}

func getDayData() (minDay int, maxDay int) {
	dateMap := model.GetDateConfig()
	dayMap := dateMap["day"].(map[string]interface{})
	minDay = int(dayMap["minday"].(float64))
	maxDay = int(dayMap["maxday"].(float64))
	return
}
