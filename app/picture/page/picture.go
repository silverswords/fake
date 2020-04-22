package page

import (
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/pkg/model"
)

//Picture is to run picture
func Picture(router *gin.Engine) {
	router.GET("/picture", picture)
}

func picture(c *gin.Context) {
	pictures := getPicture()

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(pictures))

	picturePath := model.PictureCon + pictures[n]

	c.File(picturePath)
}

func getPicture() (pictures []string) {
	dir, err := ioutil.ReadDir(model.PictureCon)
	if err != nil {
		log.Println()
	}

	for _, pi := range dir {
		pictures = append(pictures, pi.Name())
	}

	return
}
