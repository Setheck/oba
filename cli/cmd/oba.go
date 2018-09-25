package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var baseUrl string
var apiKey string

func init() {
	rootCmd.AddCommand(
		agencyCmd, blockCmd, reportCmd, routeCmd, stopCmd, tripCmd)

	if v := viper.Get("baseUrl"); v != nil {
		baseUrl = v.(string)
	}
	if v := viper.Get("apiKey"); v != nil {
		apiKey = v.(string)
	}
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
