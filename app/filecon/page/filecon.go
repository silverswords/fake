package page

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/filecon/page/model"
)

//FileCon is a router
func FileCon(router *gin.Engine) {
	router.POST("/filecon", filecon)
}

func filecon(c *gin.Context) {
	c.File(model.FilePath)
}
