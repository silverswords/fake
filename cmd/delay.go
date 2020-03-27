package cmd

import (
	"log"

	delay "github.com/silverswords/fake/app/delay"
	"github.com/spf13/cobra"
)

var delayCmd = &cobra.Command{
	Use:   "delay",
	Short: "a fake server to delay response",
	Long: `run 
	go run main.go delay`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		delay.Delay()
	},
}

func init() {
	rootCmd.AddCommand(delayCmd)
}
