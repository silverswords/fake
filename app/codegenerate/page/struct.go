package page

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
)

type inputStruct struct {
	PackageName string
	StructName  string
	Fields      []string
}

//StructGenerate is to generate struct code
func StructGenerate(router *gin.Engine) {
	router.GET("/struct", structGenerate)
}

func structGenerate(c *gin.Context) {
	codeStructMap, err := file.GetFileInterface(model.FileCodeGenerate)
	if err != nil {
		log.Println(err)
		return
	}

	data := &inputStruct{}
	data.PackageName = codeStructMap["PackageName"].(string)
	data.StructName = codeStructMap["StructName"].(string)
	fieldsMap := codeStructMap["Fields"].(map[string]interface{})
	for k := range fieldsMap {
		data.Fields = append(data.Fields, k)
	}

	shallowCopy(data, c)
}

func shallowCopy(data *inputStruct, c *gin.Context) {
	b := &bytes.Buffer{}

	fmt.Fprintf(b, "package %s\n\n", data.PackageName)
	fmt.Fprintf(b, "func (o %[1]s) ShallowCopy() %[1]s {\n", data.StructName)
	fmt.Fprintf(b, "\treturn %s{\n", data.StructName)

	for _, field := range data.Fields {
		fmt.Fprintf(b, "\t\t%[1]s: o.%[1]s,\n", field)
	}

	fmt.Fprint(b, "\t}\n")
	fmt.Fprint(b, "}\n")

	c.String(200, b.String())
}
