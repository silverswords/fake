package ws

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/silverswords/fake/wsrouter/midware"
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
	router.Use(midware.Clients())

	router.GET("/ws", func(c *gin.Context) {
		router(c.Writer, c.Request, c)
	})
}
