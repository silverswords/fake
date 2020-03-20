package page

import (
	"log"
	"math/rand"
	"net/http"

	file "github.com/silverswords/fake/pkg/file"
	model "github.com/silverswords/fake/pkg/model"
)

//Ptest is probability test
func Ptest(w http.ResponseWriter) {
	flag := rand.Intn(100)

	m, _ := file.GetFileInt(model.FileProPath)

	switch {
	case flag < m["404"]:
		log.Println("1----------------")
		http.Error(w, "Not Found", 404)
		return
	case flag >= m["404"] && flag < m["404"]+m["505"]:
		log.Println("2----------------")
		http.Error(w, "HTTP Version not supported", 505)
		return
	case flag >= m["404"]+m["505"] && flag < m["404"]+m["505"]+m["400"]:
		log.Println("3----------------")
		http.Error(w, "Bad Request", 400)
		return
	default:
		return
	}
}
