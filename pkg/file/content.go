package content

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

//GetFileInterface to get file content
func GetFileInterface(filePath string) (map[string]interface{}, error) {
	var jsonData map[string]interface{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("open error")
	}
	defer file.Close()

	jsoned, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panicln(err)
		return nil, errors.New("read error")
	}

	err = json.Unmarshal(jsoned, &jsonData)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unmarshal error")
	}

	return jsonData, nil
}

//GetFileInt to get file content
func GetFileInt(filePath string) (map[string]int, error) {
	var jsonData map[string]int

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("open error")
	}
	defer file.Close()

	jsoned, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panicln(err)
		return nil, errors.New("read error")
	}

	err = json.Unmarshal(jsoned, &jsonData)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unmarshal error")
	}

	return jsonData, nil
}
