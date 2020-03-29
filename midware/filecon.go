package midware

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/pkg/model"
)

//FileCon is to response file content
func FileCon() gin.HandlerFunc {
	return filecon
}

func filecon(c *gin.Context) {
	c.File(model.FilePath)
}
