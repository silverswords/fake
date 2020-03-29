package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/protest/midware"
	"github.com/silverswords/fake/app/protest/page"
	ws "github.com/silverswords/fake/wsconfig/ws"
)

///PreRouter is to run pre router for Upgrade
func PreRouter() {
	router := gin.Default()

	Upgrade(router)

	router.Run(":8001")
}

//Upgrade is to do router
func Upgrade(router *gin.Engine) {
	router.Use(midware.Clients())

	router.GET("/ws", func(c *gin.Context) {
		Router(c.Writer, c.Request)
	})
}

//Router is to do upgrade and do other things
func Router(w http.ResponseWriter, r *http.Request) {
	c, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	for {
		page.Ptest(c, w)
		return
	}
}
