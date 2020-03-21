package page

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	file "github.com/silverswords/fake/pkg/file"
	model "github.com/silverswords/fake/pkg/model"
)

//Ptest is probability test
func Ptest(w http.ResponseWriter, c *gin.Context) {
	flag := rand.Intn(100)

	m, _ := file.GetFileInt(model.FileProPath)

	switch {
	case flag < m["404"]:
		http.Error(w, "Not Found", 404)
		return
	case flag >= m["404"] && flag < m["404"]+m["505"]:
		http.Error(w, "HTTP Version not supported", 505)
		return
	case flag >= m["404"]+m["505"] && flag < m["404"]+m["505"]+m["400"]:
		http.Error(w, "Bad Request", 400)
		return
	default:
		return
	}
}
