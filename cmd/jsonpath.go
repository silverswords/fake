package cmd

import (
	"log"

	content "github.com/silverswords/fake/pkg/file"
	"github.com/silverswords/fake/pkg/model"
	router "github.com/silverswords/fake/router"
	"github.com/spf13/cobra"
)

var jsonpathCmd = &cobra.Command{
	Use:   "jsonpath",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		jsonData, err := content.GetFileInterface(model.FileMultiPath)
		if err != nil {
			log.Panicln(err)
			return
		}

		router.Router(jsonData)
	},
}

func init() {
	rootCmd.AddCommand(jsonpathCmd)
}
