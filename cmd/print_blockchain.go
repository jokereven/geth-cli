package cmd

import (
	"github.com/spf13/cobra"
)

var printBlockChainCmd = &cobra.Command{
	Use:   "print",
	Short: "打印区块链",
	Long:  "Print Block Chain",
	Run:   func(cmd *cobra.Command, args []string) {},
}
