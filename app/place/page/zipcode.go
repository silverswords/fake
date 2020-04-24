package page

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
)

//ZipCode is to run zipcode
func ZipCode(router *gin.Engine) {
	router.GET("/zipcode", zipCode)
}

func zipCode(c *gin.Context) {
	c.String(200, "zip code: %s", GetZipCode())
}

//GetZipCode is to get zip code
func GetZipCode() string {
	return strconv.Itoa(rand.GetRand().RandRange(100000, 999999))
}
