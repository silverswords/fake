package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

//定义空接口接收解析后的json数据
var jsonInterface map[string]interface{}

func main() {
	file, err := os.Open("/Users/lovae/go/src/github.com/silverswords/fake/test.json")
	if err != nil {
		log.Printf("open error: %s", err)
		return
	}
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(jsonData, &jsonInterface)
	if err != nil {
		log.Println(err)
		return
	}

	router := gin.Default()

	for k, v := range jsonInterface {
		va, _ := json.Marshal(v)

		router.GET("/"+k, func(cx *gin.Context) {
			cx.String(200, string(va))
		})
	}

	router.Run(":8000")
}
