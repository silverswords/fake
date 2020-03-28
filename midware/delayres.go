package midware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
)

//Delay is to do delay response
func Delay() gin.HandlerFunc {
	return delay
}

func delay(c *gin.Context) {
	time.Sleep(getTime() * time.Second)
	c.String(200, "delay succeed")
}

func getTime() time.Duration {
	delay, _ := file.GetFileInt(model.FileDelayPath)

	return time.Duration(delay["delaytime"])
}
