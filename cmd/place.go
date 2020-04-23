package cmd

import (
	"log"

	"github.com/silverswords/fake/app/place"
	"github.com/spf13/cobra"
)

var placeCmd = &cobra.Command{
	Use:   "place",
	Short: "a fake server to get router about place ",
	Long: `run 
	go run main.go place`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		place.Router()
	},
}

func init() {
	rootCmd.AddCommand(placeCmd)
}
