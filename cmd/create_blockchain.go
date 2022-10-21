package cmd

import (
	"github.com/jokereven/geth-cli/internal"
	"github.com/spf13/cobra"
)

var address string

var createBlockChainCmd = &cobra.Command{
	Use:   "create",
	Short: "创建区块链",
	Long:  "Create One new Block Chain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var createBlockChainByAddressCmd = &cobra.Command{
	Use:   "address",
	Short: "创建区块链",
	Long:  "Create One new Block Chain",
	Run: func(cmd *cobra.Command, args []string) {
		internal.TocreateBlockChain(address)
	},
}

func init() {
	createBlockChainCmd.AddCommand(createBlockChainByAddressCmd)

	createBlockChainByAddressCmd.Flags().StringVarP(&address, "address", "a", "", `创建区块链的地址`)
}
