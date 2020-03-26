package cmd

import (
	"log"

	"github.com/silverswords/fake/app/wsclient"

	"github.com/spf13/cobra"
)

var wsclientCmd = &cobra.Command{
	Use:   "wsclient",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		wsclient.WsClient()
	},
}

func init() {
	rootCmd.AddCommand(wsclientCmd)
}
