package cmd

import (
	"log"

	"github.com/silverswords/fake/app/date"
	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "a fake server to get date information",
	Long: `run 
	go run main.go date`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		date.Router()
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
