package page

import "github.com/gin-gonic/gin"

//Place is to run place
func Place(router *gin.Engine) {
	router.GET("/place", place)
}

func place(c *gin.Context) {
	c.String(200, "city: %s, country: %s", GetCity(), GetCountryE())
}
