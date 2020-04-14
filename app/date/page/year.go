package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
	"github.com/silverswords/fake/app/date/model"
)

//Year is to run year
func Year(router *gin.Engine) {
	router.GET("/year", year)
}

func year(c *gin.Context) {
	year := GetYear()

	c.String(200, "year: %d", year)
}

//GetYear is to get year
func GetYear() int {
	var year int

	if model.IsConfig("year") {
		minYear, maxYear := getYearData()
		year = rand.GetRand().RandRange(minYear, maxYear)
	} else {
		year = rand.GetRand().RandRange(2018, 2020)
	}

	return year
}

func getYearData() (minYear int, maxYear int) {
	dateMap := model.GetDateConfig()
	yearMap := dateMap["year"].(map[string]interface{})
	minYear = int(yearMap["minyear"].(float64))
	maxYear = int(yearMap["maxyear"].(float64))
	return
}
