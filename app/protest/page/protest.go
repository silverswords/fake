package page

import (
	"math/rand"

	"github.com/gorilla/websocket"
	file "github.com/silverswords/fake/pkg/file"
	model "github.com/silverswords/fake/pkg/model"
)

//Ptest is probability test
func Ptest(c *websocket.Conn) {
	flag := rand.Intn(100)

	m, _ := file.GetFileInt(model.FileProPath)

	switch {
	case flag < m["404"]:
		c.WriteMessage(websocket.TextMessage, []byte("Not Found"))
		return
	case flag >= m["404"] && flag < m["404"]+m["505"]:
		c.WriteMessage(websocket.TextMessage, []byte("HTTP Version not supported"))
		return
	case flag >= m["404"]+m["505"] && flag < m["404"]+m["505"]+m["400"]:
		c.WriteMessage(websocket.TextMessage, []byte("Bad Request"))
		return
	default:
		c.WriteMessage(websocket.TextMessage, []byte("succeed"))
		return
	}
}
