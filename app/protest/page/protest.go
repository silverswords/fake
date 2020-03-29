package page

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	file "github.com/silverswords/fake/pkg/file"
	model "github.com/silverswords/fake/pkg/model"
)

//Ptest is probability test with multi response
func Ptest(c *websocket.Conn, w http.ResponseWriter) {
	rand.Seed(time.Now().UnixNano())
	flag := rand.Intn(100)

	m, _ := file.GetFileInt(model.FileProPath)

	switch {
	case flag < m["404"]:
		http.Error(w, "Not Found", 404)
		return
	case flag >= m["404"] && flag < m["404"]+m["505"]:
		http.Error(w, "HTTP Version not supported", 505)
		return
	case flag >= m["404"]+m["505"] && flag < m["404"]+m["505"]+m["400"]:
		http.Error(w, "Bad Request", 400)
		return
	default:
		c.WriteMessage(websocket.TextMessage, []byte("succeed"))
		return
	}
}
