package main

import (
	content "github.com/silverswords/fake/content"
	errcheck "github.com/silverswords/fake/err"
	router "github.com/silverswords/fake/router"
)

func main() {
	jsonData, err := content.GetFileContent(filePath)
	errcheck.Err(err)

	router.Router(jsonData)
}
