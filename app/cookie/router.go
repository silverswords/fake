package cookie

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/cookie/page"
)

//Cookie is to run check cookie
func Cookie() {
	router := gin.Default()

	page.Check(router)

	router.Run(":8003")
}
