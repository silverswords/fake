package model

//CountrySlice is country slice
type CountrySlice []CountryStruct

//CountryStruct is country struct
type CountryStruct struct {
	Code string `json:"code"`
	En   string `json:"en"`
	Cn   string `json:"cn"`
}
