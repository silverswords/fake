package ws

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

//Upgrader is websocket upgrader
var Upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
