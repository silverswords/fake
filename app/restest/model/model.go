package model

import (
	"github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
)

//IsTest is to judge test
func IsTest() bool {
	restestMap, _ := file.GetFileBool(model.FileCodeisTest)

	return restestMap["istest"]
}

//IsLogResCode is to judge log response code
func IsLogResCode() bool {
	restestMap, _ := file.GetFileBool(model.FileCodeisTest)

	return restestMap["islogrescode"]
}
