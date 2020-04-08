package cmd

import (
	"log"

	"github.com/silverswords/fake/app/codegenerate"
	"github.com/spf13/cobra"
)

var codegenerateCmd = &cobra.Command{
	Use:   "codegenerate",
	Short: "a fake server to get router about code generate",
	Long: `run 
	go run main.go codegenerate`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		codegenerate.Router()
	},
}

func init() {
	rootCmd.AddCommand(codegenerateCmd)
}
