package file

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

//ToSlice is to get interface key to slice
func ToSlice(thisMap map[string]interface{}) (thisSlice []string) {
	for k := range thisMap {
		thisSlice = append(thisSlice, k)
	}

	return thisSlice
}
