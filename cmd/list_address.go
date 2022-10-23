package cmd

import (
	"github.com/jokereven/geth-cli/internal"
	"github.com/spf13/cobra"
)

var listAddressCmd = &cobra.Command{
	Use:   "listAddress",
	Short: "List Address",
	Long:  "List all Wallet Address",
	Run: func(cmd *cobra.Command, args []string) {
		internal.ListAddress()
	},
}
