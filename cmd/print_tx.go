package cmd

import (
	"github.com/jokereven/geth-cli/internal"
	"github.com/spf13/cobra"
)

var printTxCmd = &cobra.Command{
	Use:   "printTx",
	Short: "打印区块的所有交易",
	Long:  "Print All Block Chain Deal",
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrintTx()
	},
}
