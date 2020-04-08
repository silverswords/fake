package restest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/restest/model"
	"github.com/silverswords/fake/app/restest/page"
	"github.com/silverswords/fake/midware"
	ws "github.com/silverswords/fake/wsconfig/ws"
)

//PreRouter is to run pre router
func PreRouter() {
	router := gin.Default()

	wsconfig(router)

	router.Run(":8001")
}

//wsconfig is to add midware and do router for ws
func wsconfig(router *gin.Engine) {
	if model.IsLogResCode() {
		router.Use(midware.CodeLogger())
	}

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
		if model.IsTest() {
			page.MultiTest(c, w)
		} else {
			page.Suctest(c, w)
		}
		return
	}
}
