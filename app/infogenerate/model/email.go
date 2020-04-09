package model

import (
	"math/rand"
	"time"
)

//GetEmail is to get email
func GetEmail(number int) []string {
	var email []string

	bytes := []byte(str)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < number; i++ {
		result := []byte{}

		for j := 0; j < 10; j++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}

		email = append(email, string(result)+layout)
	}

	return email
}

var str string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var layout string = "@gmail.com"
