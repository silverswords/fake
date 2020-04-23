package page

import (
	"github.com/gin-gonic/gin"
	"github.com/goroom/rand"
)

//Phone is to run phone
func Phone(router *gin.Engine) {
	router.GET("/phone", phone)
}

func phone(c *gin.Context) {
	phone := rand.GetRand().Phone()

	c.String(200, phone)
}

//GetPhoneSlice is to get phone number
func GetPhoneSlice(number int) []string {
	var phone []string

	for i := 0; i < number; i++ {
		phone = append(phone, rand.GetRand().Phone())
	}

	return phone
}
