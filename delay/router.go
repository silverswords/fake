package router

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/delay/page"
)

//Delay is to delay response
func Delay() {
	router := gin.Default()

	page.Delay(router)

	router.Run(":8002")
}
