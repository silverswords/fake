package router

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/jsonpath/page"
)

//Router is to run router about json path
func Router(jsonData map[string]interface{}) {
	router := gin.Default()

	page.Home(router, jsonData)
	page.Member(router, jsonData)

	router.Run(":8000")
}
