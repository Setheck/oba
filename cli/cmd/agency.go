package cmd

import (
	"fmt"

	"github.com/Setheck/oba"
	"github.com/spf13/cobra"
)

func init() {
	agencyCmd.Flags().String("id", "", "agency id for lookup")
}

var agencyCmd = &cobra.Command{
	Use:   "agency",
	Short: "retrieve agencies",
	Long:  "get some agencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := oba.NewDefaultClientS(baseUrl, apiKey)
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		agency, err := client.Agency(id)
		if err != nil {
			return err
		}
		fmt.Println(agency.String())
		return nil
	},
}
