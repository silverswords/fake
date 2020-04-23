package page

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	infoModel "github.com/silverswords/fake/app/info/model"
	"github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
)

var ids []int
var names []string
var phones []string
var emails []string

//Generater is a generater to generate information
func Generater(router *gin.Engine) {
	router.GET("/infogenerate", generater)
}

func generater(c *gin.Context) {
	configMap, err := file.GetFileInterface(model.FileInfo)
	if err != nil {
		log.Println(err)
		return
	}

	infoMap := configMap["info"].(map[string]interface{})
	number := int(configMap["number"].(float64))

	for k := range infoMap {
		if k == "id" {
			ids = GetIDOrderSlice(number)
		} else if k == "name" {
			names = GetNameSlice(number)
		} else if k == "phone" {
			phones = GetPhoneSlice(number)
		} else if k == "email" {
			emails = GetEmailSlice(number)
		} else {
			fmt.Println("config parameter error")
			return
		}
	}

	for i := 0; i < number; i++ {
		info := &infoModel.InfoStruct{}

		for k := range infoMap {
			if k == "id" {
				info.ID = ids[i]
			} else if k == "name" {
				info.Name = names[i]
			} else if k == "phone" {
				info.Phone = phones[i]
			} else if k == "email" {
				info.Email = emails[i]
			} else {
				fmt.Println("config parameter error")
				return
			}
		}

		infoJSON, err := json.Marshal(info)
		if err != nil {
			fmt.Println(err)
			return
		}

		c.String(200, string(infoJSON)+"\n")
	}
}
