package page

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	file "github.com/silverswords/fake/pkg/file"
	model "github.com/silverswords/fake/pkg/model"
)

//MultiTest is probability test with multi response by restest in config
func MultiTest(c *websocket.Conn, w http.ResponseWriter) {
	rand.Seed(time.Now().UnixNano())
	flag := rand.Float32()

	m, _ := file.GetFilefloat32(model.FileProPath)

	var s []string
	for key := range m {
		s = append(s, key)
	}

	for i := 0; i < len(m); i++ {
		if flag < m[s[i]] {
			rescode, _ := strconv.Atoi(s[i])
			http.Error(w, s[i], rescode)
			return
		}

		m[s[i+1]] = m[s[i+1]] + m[s[i]]
	}
}
