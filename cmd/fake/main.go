package main

import (
	"log"

	content "github.com/silverswords/fake/pkg/file"
	model "github.com/silverswords/fake/pkg/model"
	router "github.com/silverswords/fake/router"
)

func main() {
	jsonData, err := content.GetFileContent(model.FilePath)
	if err != nil {
		log.Panicln(err)
		return
	}

	router.Router(jsonData)
}
