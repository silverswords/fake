package router

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/jsonpath/page"
)

//Router is to create router
func Router(jsonData map[string]interface{}) {
	router := gin.Default()

	page.Home(router, jsonData)
	page.JSONCon(router, jsonData)

	router.Run(":8000")
}
