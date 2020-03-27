package wsclient

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

//WsClient is to run a websocket client
func WsClient() {
	u := url.URL{Scheme: "ws", Host: "localhost:8001", Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		return
	}

}
