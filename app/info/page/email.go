package page

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/info/model"
)

//Email is to run email
func Email(router *gin.Engine) {
	router.GET("/email", email)
}

func email(c *gin.Context) {
	email := GetEmail()

	c.String(200, email)
}

//GetEmail is to get email
func GetEmail() string {
	var email string

	bytes := []byte(model.Characters)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := []byte{}

	for j := 0; j < 10; j++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	email = string(result) + model.Gmail

	return email
}

//GetEmailSlice is to get email slice
func GetEmailSlice(number int) []string {
	var emails []string

	for i := 0; i < number; i++ {
		emails = append(emails, GetEmail())
	}

	return emails
}
