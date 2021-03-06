package page

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/place/model"
)

//City is to run city
func City(router *gin.Engine) {
	router.GET("/city", city)
}

//city is to get city randomly
func city(c *gin.Context) {
	city := GetCity()

	c.String(200, city)
}

//GetCity is to get a city about china
func GetCity() string {
	var city string

	rand.Seed(time.Now().UnixNano())

	city = model.CitySlice[rand.Intn(model.CityLenth)]

	return city
}
