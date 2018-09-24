package cmd

import (
	"fmt"
	"github.com/Setheck/oba"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	routeCmd.Flags().String("id", "", "route id for lookup")
}

var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "retrieve routes",
	Long:  "get some routes",
	RunE: func(cmd *cobra.Command, args []string) error {
		baseUrl := viper.Get("baseUrl").(string)
		apiKey := viper.Get("apiKey").(string)
		client := oba.NewDefaultClientS(baseUrl, apiKey)
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		route, err := client.Route(id)
		if err != nil {
			return err
		}
		fmt.Println(route.String())
		return nil
	},
}
