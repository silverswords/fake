package cmd

import (
	"log"

	"github.com/silverswords/fake/app/wsclient"

	"github.com/spf13/cobra"
)

var wsclientCmd = &cobra.Command{
	Use:   "wsclient",
	Short: "a fake server to run websockte client",
	Long: `run 
	go run main.go wsclient`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		wsclient.WsClient()
	},
}

func init() {
	rootCmd.AddCommand(wsclientCmd)
}
