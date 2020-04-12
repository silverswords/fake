package cmd

import (
	"log"

	"github.com/silverswords/fake/app/city"
	"github.com/spf13/cobra"
)

var cityCmd = &cobra.Command{
	Use:   "city",
	Short: "a fake server to get router about city of China ",
	Long: `run 
	go run main.go city`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		city.Router()
	},
}

func init() {
	rootCmd.AddCommand(cityCmd)
}
