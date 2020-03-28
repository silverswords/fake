package midware

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/pkg/file"
	model "github.com/silverswords/fake/pkg/model"
)

//Check middleware is to check if specific cookies are present
func Check() gin.HandlerFunc {
	return check
}

func check(c *gin.Context) {
	cookie, err := c.Cookie(getCookie())

	if err != nil {
		c.String(200, "no cookie:%s", err)
	}

	c.String(200, "Cookie value: %s \n", cookie)
}

func getCookie() string {
	cookie, _ := file.GetFileInterface(model.FileCookiePath)
	return cookie["cookie"].(string)
}
