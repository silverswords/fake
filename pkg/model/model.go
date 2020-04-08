package model

var (
	//FileMultiPath is the path of multi
	FileMultiPath = "/Users/lovae/go/src/github.com/silverswords/fake/config/multi.json"
	//FileRestestPath is the path of pro
	FileRestestPath = "/Users/lovae/go/src/github.com/silverswords/fake/config/restest.json"
	//FileDelayPath is the path of delay time
	FileDelayPath = "/Users/lovae/go/src/github.com/silverswords/fake/config/delay.json"
	//FileCookiePath is the path of cookie
	FileCookiePath = "/Users/lovae/go/src/github.com/silverswords/fake/config/cookie.json"
	//FilePath is file path for response file content
	FilePath = "/Users/lovae/go/src/github.com/silverswords/fake/config/file.txt"
	// FileCodePath is the file path of response file
	FileCodePath = "/Users/lovae/go/src/github.com/silverswords/fake/code.json"
	// FileIsRestestTest is the file path of istest
	FileIsRestestTest = "/Users/lovae/go/src/github.com/silverswords/fake/config/isrestest.json"
	// FileCodeGenerate is the file path of code generate
	FileCodeGenerate = "/Users/lovae/go/src/github.com/silverswords/fake/config/codestruct.json"
)

//InFor is the information of bode
type InFor struct {
	ID string `form:"id" json:"id" xml:"id"  binding:"required"`
}

//InfoHeader is the information of header
type InfoHeader struct {
	Function string `header:"function" binding:"required"`
}
