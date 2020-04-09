package model

import (
	"github.com/goroom/rand"
)

//GetPhone is to get phone number
func GetPhone(number int) []string {
	var phone []string

	for i := 0; i < number; i++ {
		phone = append(phone, rand.GetRand().Phone())
	}

	return phone
}
