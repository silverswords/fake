package main

import "github.com/silverswords/fake/wsrouter"

func main() {
	/*
		jsonData, err := content.GetFileContent(model.FilePath)
		if err != nil {
			log.Panicln(err)
			return
		}

		router.Router(jsonData)
	*/

	wsrouter.WsRouter()
}
