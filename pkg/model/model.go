package model

var (
	//FileMultiPath is the path of multi
	FileMultiPath = fileRoot + "github.com/silverswords/fake/config/jsonpath/multi.json"
	//FileRestestPath is the path of restest
	FileRestestPath = fileRoot + "github.com/silverswords/fake/config/restest/restest.json"
	// FileCodeGenerate is the file path of code generate
	FileCodeGenerate = fileRoot + "github.com/silverswords/fake/config/code/codestruct.json"
	//FileInfo is the file path of info
	FileInfo = fileRoot + "github.com/silverswords/fake/config/info/info/info.json"
	//FileDate is the file path of date
	FileDate = fileRoot + "github.com/silverswords/fake/config/date/date.json"
	//PictureCon is the path of picture config
	PictureCon = fileRoot + "github.com/silverswords/fake/config/picture/"
	// FileCountryPath is the path of country
	FileCountryPath = fileRoot + "github.com/yimkh/country/country.json"
)

//about midware
var (
	//FilePath is file path for response file content
	FilePath = fileRoot + "github.com/silverswords/fake/config/midware/file.txt"
	//FileDelayPath is the path of delay time
	FileDelayPath = fileRoot + "github.com/silverswords/fake/config/midware/delay.json"
	//FileCookiePath is the path of cookie
	FileCookiePath = fileRoot + "github.com/silverswords/fake/config/midware/cookie.json"
)

//need to modify your file path
var fileRoot = "/Users/lovae/go/src/"

//InFor is the information of bode
type InFor struct {
	ID string `form:"id" json:"id" xml:"id"  binding:"required"`
}

//InfoHeader is the information of header
type InfoHeader struct {
	Function string `header:"function" binding:"required"`
}
