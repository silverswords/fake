package cmd

import (
	"log"

	protest "github.com/silverswords/fake/app/protest"
	"github.com/spf13/cobra"
)

var protestCmd = &cobra.Command{
	Use:   "protest",
	Short: "a fake server to multi return with websocket",
	Long: `run 
	go run main.go protest`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		protest.PreRouter()
	},
}

func init() {
	rootCmd.AddCommand(protestCmd)
}
