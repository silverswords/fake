package cmd

import (
	"log"

	restest "github.com/silverswords/fake/app/restest"
	"github.com/spf13/cobra"
)

var restestCmd = &cobra.Command{
	Use:   "restest",
	Short: "a fake server to multi return with websocket",
	Long: `run 
	go run main.go restest`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		restest.PreRouter()
	},
}

func init() {
	rootCmd.AddCommand(restestCmd)
}
