package cmd

import (
	"log"

	"github.com/silverswords/fake/app/infogenerate"
	"github.com/spf13/cobra"
)

var infogenerateCmd = &cobra.Command{
	Use:   "infogenerate",
	Short: "a fake server to get router about information generate",
	Long: `run 
	go run main.go infogenerate`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		infogenerate.Router()
	},
}

func init() {
	rootCmd.AddCommand(infogenerateCmd)
}
