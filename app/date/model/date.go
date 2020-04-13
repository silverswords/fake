package model

//DateStruct is data struct
type DateStruct struct {
	Year   int    `json:"year"`
	Month  string `json:"month"`
	Day    int    `json:"day"`
	Hour   int    `json:"hour"`
	Second int    `json:"second"`
}
