package cmd

import (
	"log"

	"github.com/silverswords/fake/app/commerce"
	"github.com/spf13/cobra"
)

var commerceCmd = &cobra.Command{
	Use:   "commerce",
	Short: "a fake server to get commerce information",
	Long: `run 
	go run main.go commerce`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		commerce.Router()
	},
}

func init() {
	rootCmd.AddCommand(commerceCmd)
}
