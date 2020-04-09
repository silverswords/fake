package model

//InfoStruct is the information struct
type InfoStruct struct {
	ID    int    `json:"id"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}
