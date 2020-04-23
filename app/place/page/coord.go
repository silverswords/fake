package page

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

//Coord is to run coord
func Coord(router *gin.Engine) {
	router.GET("/coord", coord)
}

func coord(c *gin.Context) {
	c.String(200, "longitude: %f, latitude: %f", GetLongitude(), GetLatitude())
}

//GetLongitude is to get longitude
func GetLongitude() float32 {
	rand.Seed(time.Now().UnixNano())

	longitude := rand.Float32()*360 - 180

	return longitude
}

//GetLatitude is to get latitude
func GetLatitude() float32 {
	rand.Seed(time.Now().UnixNano())

	latitude := rand.Float32()*180 - 90

	return latitude
}
