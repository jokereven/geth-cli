package cmd

import (
	"log"

	"github.com/jokereven/geth-cli/internal"

	"github.com/spf13/cobra"
)

var createWalletCmd = &cobra.Command{
	Use:   "createWallet",
	Short: "创建钱包",
	Long:  "Create One Wallet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var toCreateWalletCmd = &cobra.Command{
	Use:   "new",
	Short: "创建钱包",
	Long:  "Create One Wallet",
	Run: func(cmd *cobra.Command, args []string) {
		internal.CreateWallet()
		log.Println("Create One New Wallet")
	},
}

func init() {
	createWalletCmd.AddCommand(toCreateWalletCmd)
}
