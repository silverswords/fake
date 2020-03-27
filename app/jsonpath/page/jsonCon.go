package page

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

//JSONCon is to response json path
func JSONCon(router *gin.Engine, jsonData map[string]interface{}) {
	for k, v := range jsonData {
		value, _ := json.Marshal(v)
		router.GET("/"+k, func(c *gin.Context) {
			c.String(200, string(value))
		})
	}
}
