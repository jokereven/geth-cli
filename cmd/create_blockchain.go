package cmd

import (
	"github.com/spf13/cobra"
)

var createBlockChainCmd = &cobra.Command{
	Use:   "create",
	Short: "创建区块链",
	Long:  "Create One new Block Chain",
	Run:   func(cmd *cobra.Command, args []string) {},
}
