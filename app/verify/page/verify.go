package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/silverswords/fake/app/verify/medel"
)

//Verify is to do Verifying
func Verify(router *gin.Engine) {
	router.POST("/info", info)
}

func info(c *gin.Context) {
	infoheader := model.InfoHeader{}

	if err := c.ShouldBindHeader(&infoheader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	infor := model.InFor{}

	if err := c.ShouldBind(&infor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if infoheader.Function != "verify" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unsucceed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inFor": infor.ID})
}
