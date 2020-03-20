package wsrouter

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/wsrouter/ws"
)

//WsRouter is to upgrade websocket
func WsRouter() {
	router := gin.Default()

	ws.Upgrade(router)

	router.Run(":8001")
}
