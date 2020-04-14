package model

import (
	"log"

	"github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
)

//DateStruct is data struct
type DateStruct struct {
	Year   int    `json:"year"`
	Month  string `json:"month"`
	Day    int    `json:"day"`
	Hour   int    `json:"hour"`
	Minute int    `json:"minute"`
	Second int    `json:"second"`
}

//GetDateConfig is to get date config
func GetDateConfig() map[string]interface{} {
	dateMap, err := file.GetFileInterface(model.FileDate)
	if err != nil {
		log.Println(err)
	}

	return dateMap
}

//IsConfig is to judge wether config is exist
func IsConfig(name string) bool {
	dateMap := GetDateConfig()

	for k := range dateMap {
		if k == name {
			return true
		}
	}
	return false
}
