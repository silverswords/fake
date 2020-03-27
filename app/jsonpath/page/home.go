package page

import "github.com/gin-gonic/gin"

var (
	rootDirectly = "localhost:8000/"
)

//Home is to response home path
func Home(router *gin.Engine, jsonData map[string]interface{}) {
	router.GET("/home", func(c *gin.Context) {
		for k := range jsonData {
			c.String(200, rootDirectly+k+"\n")
		}
	})
}
