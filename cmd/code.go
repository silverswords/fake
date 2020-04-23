package cmd

import (
	"log"

	"github.com/silverswords/fake/app/code"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "a fake server to get router about code generate",
	Long: `run 
	go run main.go code`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		code.Router()
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)
}
