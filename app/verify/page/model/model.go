package model

//InFor is the information of bode
type InFor struct {
	ID string `form:"id" json:"id" xml:"id"  binding:"required"`
}

//InfoHeader is the information of header
type InfoHeader struct {
	Function string `header:"function" binding:"required"`
}
