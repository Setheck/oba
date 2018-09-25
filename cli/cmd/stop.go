package cmd

import (
	"fmt"

	"github.com/Setheck/oba"
	"github.com/spf13/cobra"
)

func init() {
	stopCmd.Flags().String("id", "", "stop id for lookup")
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "retrieve stops",
	Long:  "get some stops",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := oba.NewDefaultClientS(baseUrl, apiKey)
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		stop, err := client.Stop(id)
		if err != nil {
			return err
		}
		fmt.Println(stop)
		return nil
	},
}
