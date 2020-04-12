package page

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/date/model"
)

//Date is to run data
func Date(router *gin.Engine) {
	router.GET("/date", date)
}

func date(c *gin.Context) {
	var date model.DateStruct

	date.Year = GetYear()
	date.Month = GetMonth()

	c.JSON(200, date)
}
