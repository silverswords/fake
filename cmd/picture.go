package cmd

import (
	"log"

	"github.com/silverswords/fake/app/picture"
	"github.com/spf13/cobra"
)

var pictureCmd = &cobra.Command{
	Use:   "picture",
	Short: "a fake server to get router about picture ",
	Long: `run 
	go run main.go picture`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		picture.Router()
	},
}

func init() {
	rootCmd.AddCommand(pictureCmd)
}
