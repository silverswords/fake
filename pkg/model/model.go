package model

var (
	//FileMultiPath is the path of multi
	FileMultiPath = fileRoot + "github.com/silverswords/fake/config/multi.json"
	//FileRestestPath is the path of pro
	FileRestestPath = fileRoot + "github.com/silverswords/fake/config/restest.json"
	//FileDelayPath is the path of delay time
	FileDelayPath = fileRoot + "github.com/silverswords/fake/config/delay.json"
	//FileCookiePath is the path of cookie
	FileCookiePath = fileRoot + "github.com/silverswords/fake/config/cookie.json"
	//FilePath is file path for response file content
	FilePath = fileRoot + "github.com/silverswords/fake/config/file.txt"
	// FileCodePath is the file path of response file
	FileCodePath = fileRoot + "github.com/silverswords/fake/code.json"
	// FileIsRestestTest is the file path of istest
	FileIsRestestTest = fileRoot + "github.com/silverswords/fake/config/isrestest.json"
	// FileCodeGenerate is the file path of code generate
	FileCodeGenerate = fileRoot + "github.com/silverswords/fake/config/codestruct.json"
	//FileInfo is the file path of info
	FileInfo = fileRoot + "github.com/silverswords/fake/config/info.json"
	//FileDate is the file path of date
	FileDate = fileRoot + "github.com/silverswords/fake/config/date.json"
	//PictureCon is the path of picture config
	PictureCon = fileRoot + "github.com/silverswords/fake/config/picture/"
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
