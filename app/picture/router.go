package picture

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/picture/page"
)

//Router is to run Picture
func Router() {
	router := gin.Default()

	page.Picture(router)

	router.Run(":8006")
}
