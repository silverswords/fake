package midware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/pkg/file"
	model "github.com/silverswords/fake/pkg/model"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

//CodeLogger is to log the number of response code
func CodeLogger() gin.HandlerFunc {
	return ginBodyLogMiddleware
}

func ginBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()

	statisticsMap := preMap()

	printStatisticsMap(statisticsMap, statusCode)

	fmt.Print("Response body: " + blw.body.String())
}

func preMap() map[string]float32 {
	m, _ := file.GetFileInterface(model.FileRestestPath)

	preMap := make(map[string]float32, len(m))

	for key := range m {
		preMap[key] = 0
	}

	return preMap
}

func printStatisticsMap(statisticsMap map[string]float32, statusCode int) {
	for key := range statisticsMap {
		if strconv.Itoa(statusCode) == key {
			statisticsMap[key]++
		}
	}

	statisticsJSON, err := json.Marshal(statisticsMap)
	if err != nil {
		log.Println(err)
		return
	}

	codejsonfile, err := os.OpenFile("code.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer codejsonfile.Close()

	preMap, err := file.GetFileInterface("code.json")

	m := make(map[string]float64, len(preMap))

	for k, v := range preMap {
		m[k] = v.(float64)
	}

	if err != nil {
		log.Println(err)
	}
	if m == nil {
		writer := bufio.NewWriter(codejsonfile)
		writer.Write(statisticsJSON)
		writer.Flush()
	} else {
		for key := range m {
			if strconv.Itoa(statusCode) == key {
				m[key]++
				break
			}
		}

		statisticsJSON, err = json.Marshal(m)
		if err != nil {
			log.Println(err)
			return
		}

		writer := bufio.NewWriter(codejsonfile)
		writer.Write(statisticsJSON)
		writer.Flush()
	}
}
