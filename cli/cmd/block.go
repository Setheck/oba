package cmd

import (
	"fmt"

	"github.com/Setheck/oba"
	"github.com/spf13/cobra"
)

func init() {
	blockCmd.Flags().String("id", "", "block id for lookup")
}

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "retrieve blocks",
	Long:  "get some blocks",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := oba.NewDefaultClientS(baseUrl, apiKey)
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		block, err := client.Block(id)
		if err != nil {
			return err
		}
		fmt.Println(block)
		return nil
	},
}
