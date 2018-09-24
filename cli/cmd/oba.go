package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(routeCmd)
}

var rootCmd = &cobra.Command{
	Use:   "oba",
	Short: "OneBusAway Cli tool",
	Long:  "OBA Cli tool!",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
