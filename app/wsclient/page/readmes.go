package readmes

import (
	"log"

	"github.com/gorilla/websocket"
)

//Readmes is to read message
func Readmes(c *websocket.Conn) {
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", message)
}
