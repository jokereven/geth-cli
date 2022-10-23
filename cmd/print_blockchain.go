package cmd

import (
	"github.com/jokereven/geth-cli/internal"
	"github.com/spf13/cobra"
)

var printBlockChainCmd = &cobra.Command{
	Use:   "print",
	Short: "Print Block Chain",
	Long:  "Print Block Chain",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Print()
	},
}
