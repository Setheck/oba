package cmd

import (
	"fmt"

	"github.com/Setheck/oba"
	"github.com/spf13/cobra"
)

func init() {
	agencyCmd.Flags().String("id", "", "agency id for lookup")
	agencyCmd.Flags().Bool("coverage", false, "agencies with coverage")
}

var agencyCmd = &cobra.Command{
	Use:   "agency",
	Short: "retrieve agencies",
	Long:  "get some agencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := oba.NewDefaultClientS(baseUrl, apiKey)
		coverage, err := cmd.Flags().GetBool("coverage")
		if err != nil {
			return err
		}
		if coverage {
			return AgenciesWithCoverage(client)
		}
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		return Agency(client, id)
	},
}

func Agency(client oba.Client, id string) error {
	agency, err := client.Agency(id)
	if err != nil {
		return err
	}
	fmt.Println(agency.String())
	return nil
}

func AgenciesWithCoverage(client oba.Client) error {
	awc, err := client.AgenciesWithCoverage()
	if err != nil {
		return err
	}
	for _, a := range awc {
		fmt.Println(a)
	}
	return nil
}
