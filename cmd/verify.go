package cmd

import (
	"log"

	verify "github.com/silverswords/fake/app/verify"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "a fake server to verify infomation",
	Long: `run 
	go run main.go verify`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		verify.Verify()
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
