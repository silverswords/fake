package page

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//ID is to run ID
func ID(router *gin.Engine) {
	router.GET("/id", id)
}

func id(c *gin.Context) {
	id := GetID(100)

	idst := strconv.Itoa(id)

	c.String(200, idst)
}

//GetID is to get ID
func GetID(number int) int {
	rand.Seed(time.Now().UnixNano())

	id := rand.Intn(number)

	return id
}

//GetIDOrderSlice is to get id slice in order
func GetIDOrderSlice(number int) []int {
	var id []int

	for i := 0; i < number; i++ {
		id = append(id, i)
	}

	return id
}
