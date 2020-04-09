package page

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	infoModel "github.com/silverswords/fake/app/infogenerate/model"
	"github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
)

var id []int
var name []string
var phone []string
var email []string

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
			id = infoModel.GetID(number)
		} else if k == "name" {
			name = infoModel.GetName(number)
		} else if k == "phone" {
			phone = infoModel.GetPhone(number)
		} else if k == "email" {
			email = infoModel.GetEmail(number)
		} else {
			fmt.Println("config parameter error")
			return
		}
	}

	for i := 0; i < number; i++ {
		info := &infoModel.InfoStruct{}

		for k := range infoMap {
			if k == "id" {
				info.ID = id[i]
			} else if k == "name" {
				info.Name = name[i]
			} else if k == "phone" {
				info.Phone = phone[i]
			} else if k == "email" {
				info.Email = email[i]
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
