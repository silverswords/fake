package page

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

//Member is to response json path about detail
func Member(router *gin.Engine, jsonData map[string]interface{}) {
	for k, v := range jsonData {
		value, _ := json.Marshal(v)

		router.GET("/"+k, func(c *gin.Context) {
			c.String(200, string(value))
		})
	}
}
