package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/protest/midware"
	"github.com/silverswords/fake/app/protest/page"
	ws "github.com/silverswords/fake/wsconfig/ws"
)

//PreRouter is to do Upgrade
func PreRouter() {
	router := gin.Default()

	Upgrade(router)

	router.Run(":8001")
}

//Upgrade is to upgrade
func Upgrade(router *gin.Engine) {
	router.Use(midware.Clients())

	router.GET("/protest", func(c *gin.Context) {
		Router(c.Writer, c.Request)
	})
}

//Router is to do
func Router(w http.ResponseWriter, r *http.Request) {
	_, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		page.Ptest(w)
		return
	}
}
