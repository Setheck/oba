package cmd

import (
	"fmt"

	"github.com/Setheck/oba"
	"github.com/spf13/cobra"
)

func init() {
	tripCmd.Flags().String("id", "", "trip id for lookup")
}

var tripCmd = &cobra.Command{
	Use:   "trip",
	Short: "retrieve trips",
	Long:  "get some trips",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := oba.NewDefaultClientS(baseUrl, apiKey)
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		trip, err := client.Trip(id)
		if err != nil {
			return err
		}
		fmt.Println(trip)
		return nil
	},
}
