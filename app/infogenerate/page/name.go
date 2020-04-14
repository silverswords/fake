package page

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/infogenerate/model"
)

//Name is to run name
func Name(router *gin.Engine) {
	router.GET("/name", name)
}

func name(c *gin.Context) {
	name := GetName()
	c.String(200, name)
}

//GetName is to get name
func GetName() string {
	var name string

	rand.Seed(time.Now().UnixNano())

	var first string
	var last string

	//随机产生2位或者3位的名
	for i := 0; i <= rand.Intn(2); i++ {
		first = first + model.FirstName[rand.Intn(model.FirstNameLen-1)]
	}
	last = model.LastName[rand.Intn(model.LastNameLen-1)]

	name = last + first

	return name
}

//GetNameSlice is to get name slice
func GetNameSlice(number int) []string {
	var name []string

	for i := 0; i < number; i++ {
		name = append(name, GetName())
	}

	return name
}
