package date

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/date/page"
)

//Router is to run date router
func Router() {
	router := gin.Default()

	page.Date(router)
	page.Year(router)

	router.Run(":8005")
}
