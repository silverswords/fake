package cmd

import (
	"log"

	filecon "github.com/silverswords/fake/app/filecon"
	"github.com/spf13/cobra"
)

var fileconCmd = &cobra.Command{
	Use:   "filecon",
	Short: "a fake server to response file content",
	Long: `run 
	go run main.go filecon`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		filecon.FileContent()
	},
}

func init() {
	rootCmd.AddCommand(fileconCmd)
}
