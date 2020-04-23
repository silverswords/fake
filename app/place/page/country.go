package page

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/app/place/model"
	fileModel "github.com/silverswords/fake/pkg/model"
)

//Country is to run country
func Country(router *gin.Engine) {
	router.GET("/country", country)
}

func country(c *gin.Context) {
	country := GetCountry()

	c.JSON(200, country)
}

//GetCountry is to get country
func GetCountry() model.CountryStruct {
	country := GetCountrySlice()

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(country))

	return country[n]
}

//GetCountryC is to get country about Chinese
func GetCountryC() string {
	country := GetCountrySlice()

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(country))

	return country[n].Cn
}

//GetCountryE is to get country about English
func GetCountryE() string {
	country := GetCountrySlice()

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(country))

	return country[n].En
}

//GetCountryCode is to get country code
func GetCountryCode() string {
	country := GetCountrySlice()

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(country))

	return country[n].Code
}

//GetCountrySlice is to get country information from file config
func GetCountrySlice() model.CountrySlice {
	countryFile, err := os.Open(fileModel.FileCountryPath)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer countryFile.Close()

	countryByte, err := ioutil.ReadAll(countryFile)
	if err != nil {
		log.Println(err)
		return nil
	}

	var cs model.CountrySlice

	err = json.Unmarshal(countryByte, &cs)
	if err != nil {
		log.Println(err)
		return nil
	}

	return cs
}
