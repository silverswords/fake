package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/wsrouter/page"
)

//Router is to do
func Router(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		page.Ptest(w, c)
		return
	}
}
