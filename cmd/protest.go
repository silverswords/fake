package cmd

import (
	"log"

	protest "github.com/silverswords/fake/protest"
	"github.com/spf13/cobra"
)

var protestCmd = &cobra.Command{
	Use:   "protest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting success")

		protest.PreRouter()
	},
}

func init() {
	rootCmd.AddCommand(protestCmd)
}