package cmd

import (
	"log"

	"github.com/silverswords/fake/app/info"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "a fake server to get router about information generate",
	Long: `run 
	go run main.go info`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		info.Router()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
