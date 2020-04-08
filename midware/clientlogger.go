package midware

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/pkg/file"
)

//Clients is to log clients
func Clients() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientIP := c.ClientIP()

		file.WriteFile("gin.clients", clientIP)

		return
	}
}
