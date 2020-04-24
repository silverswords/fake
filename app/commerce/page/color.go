package page

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
)

//Color is to run color
func Color(router *gin.Engine) {
	router.GET("/color", color)
}

func color(c *gin.Context) {
	c.String(200, "color: %s", GetColor())
}

//GetColor is to get color
func GetColor() string {
	colorMap, err := file.GetFileInterface(model.FileColorPath)
	if err != nil {
		log.Println(err)
		return ""
	}

	colorSlice := file.ToSlice(colorMap)

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(colorSlice))

	return colorSlice[n]
}
