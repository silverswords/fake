package cmd

import (
	"log"

	cookie "github.com/silverswords/fake/app/cookie"
	"github.com/spf13/cobra"
)

var cookieCmd = &cobra.Command{
	Use:   "cookie",
	Short: "a fake server to check cookie",
	Long: `run 
	go run main.go cookie`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		cookie.Cookie()
	},
}

func init() {
	rootCmd.AddCommand(cookieCmd)
}
