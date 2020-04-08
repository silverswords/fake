package codegenerate

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/codegenerate/page"
)

//Router is to
func Router() {
	router := gin.Default()

	page.StructGenerate(router)

	router.Run(":8002")
}
