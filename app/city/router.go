package city

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/city/page"
)

//Router is to run city router
func Router() {
	router := gin.Default()

	page.City(router)

	router.Run(":8004")
}
