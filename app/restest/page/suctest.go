package page

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Suctest is always to response 200 succeed
func Suctest(c *websocket.Conn, w http.ResponseWriter) {
	http.Error(w, "succeed", 200)
}
