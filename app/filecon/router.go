package filecon

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/filecon/page"
)

//FileContent is to run file content
func FileContent() {
	router := gin.Default()

	page.FileCon(router)

	router.Run(":8005")
}
