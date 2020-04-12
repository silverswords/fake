package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
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
	year := rand.GetRand().RandRange(2018, 2020)

	return year
}
