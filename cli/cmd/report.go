package cmd

import (
	"github.com/Setheck/oba"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "report things",
	Long:  "send a report",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := oba.NewDefaultClientS(baseUrl, apiKey)
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		return client.ReportProblemWithStop(id, nil)
	},
}
