package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/silverswords/fake/pkg/model"
)

//IsTest is to judge test
func IsTest() bool {
	resTest := GetResTest()

	isTest := resTest.Test.Istest

	return isTest
}

//IsLogResCode is to judge log response code
func IsLogResCode() bool {
	resTest := GetResTest()

	islogrescode := resTest.Test.Islogrescode

	return islogrescode
}

//GetResTest is to get restest struct
func GetResTest() ResTestStuct {
	resFile, err := os.Open(model.FileRestestPath)
	if err != nil {
		log.Println(err)
	}
	defer resFile.Close()

	resByte, err := ioutil.ReadAll(resFile)
	if err != nil {
		log.Println(err)
	}

	var rt ResTestStuct

	err = json.Unmarshal(resByte, &rt)

	return rt
}

//ResTestStuct is restest struct
type ResTestStuct struct {
	Code map[string]float32 `json:"code"`
	Test Test               `json:"test"`
}

//Test is test struct
type Test struct {
	Istest       bool `json:"istest"`
	Islogrescode bool `json:"islogrescode"`
}
