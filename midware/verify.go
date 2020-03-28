package midware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/silverswords/fake/pkg/model"
)

//Verify is to do Verifying information
func Verify() gin.HandlerFunc {
	return info
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
