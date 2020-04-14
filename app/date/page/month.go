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
	var month string

	if model.IsConfig("month") {
		month = getMonthData()
	} else {
		month = rand.GetRand().StringArray(&model.MonthData)
	}

	return month
}

func getMonthData() string {
	var month string
	var months []string

	dateMap := model.GetDateConfig()
	monthMap := dateMap["month"].(map[string]interface{})

	for k := range monthMap {
		months = append(months, k)
	}

	month = rand.GetRand().StringArray(&months)
	return month
}
