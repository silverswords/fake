package infogenerate

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/infogenerate/page"
)

//Router is to run infogenerate router
func Router() {
	router := gin.Default()

	page.Generater(router)

	router.Run(":8003")
}
