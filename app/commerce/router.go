package commerce

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/commerce/page"
)

//Router is to run color router
func Router() {
	router := gin.Default()

	page.Color(router)

	router.Run(":8007")
}
