package cmd

import (
	"fmt"

	"github.com/Setheck/oba"
	"github.com/spf13/cobra"
)

func init() {
	routeCmd.Flags().String("id", "", "route id for lookup")
	routeCmd.Flags().String("aid", "", "route ids for agency id [aid] lookup")
}

var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "retrieve routes",
	Long:  "get some routes",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := oba.NewDefaultClientS(baseUrl, apiKey)
		aid, err := cmd.Flags().GetString("aid")
		if err != nil {
			return err
		}
		return RouteIdsForAgency(client, aid)

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		return Route(client, id)
	},
}

func Route(client oba.Client, id string) error {
	route, err := client.Route(id)
	if err != nil {
		return err
	}
	fmt.Println(route.String())
	return nil
}

func RouteIdsForAgency(client oba.Client, agency string) error {
	ids, err := client.RouteIdsForAgency(agency)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", ids)
	return nil
}
