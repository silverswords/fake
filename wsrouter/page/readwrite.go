package page

import (
	"log"

	"github.com/gorilla/websocket"
)

//Read is to read message
func Read(c *websocket.Conn) {
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", message)
	log.Printf("recv type: %d", mt)
}

//Write is to write message
func Write(c *websocket.Conn, mt int, message []byte) {
	err := c.WriteMessage(mt, message)
	if err != nil {
		log.Println("write:", err)
		return
	}
}
