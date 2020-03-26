package verify

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/verify/page"
)

// Verify is to verifying infomation
func Verify() {
	router := gin.Default()

	page.Verify(router)

	router.Run(":8004")
}