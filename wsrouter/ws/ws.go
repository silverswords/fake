package ws

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//Upgrader is websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Upgrade is to upgrade
func Upgrade(router *gin.Engine) {
	router.GET("/ws", func(c *gin.Context) {
		Router(c.Writer, c.Request)
	})
}
