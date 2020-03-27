package router

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/delay/page"
)

//Delay is to run delay response
func Delay() {
	router := gin.Default()

	page.Delay(router)

	router.Run(":8002")
}
