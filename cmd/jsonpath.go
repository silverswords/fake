package cmd

import (
	"log"

	jsonpath "github.com/silverswords/fake/app/jsonpath"
	content "github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
	"github.com/spf13/cobra"
)

var jsonpathCmd = &cobra.Command{
	Use:   "jsonpath",
	Short: "a fake server to get router about json",
	Long: `run 
	go run main.go jsonpath`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		jsonData, err := content.GetFileInterface(model.FileMultiPath)
		if err != nil {
			log.Panicln(err)
			return
		}

		jsonpath.Router(jsonData)
	},
}

func init() {
	rootCmd.AddCommand(jsonpathCmd)
}
