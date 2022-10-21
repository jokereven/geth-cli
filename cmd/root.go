package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// add to root
	rootCmd.AddCommand(addBlockCmd)
	rootCmd.AddCommand(createBlockChainCmd)
	rootCmd.AddCommand(createWalletCmd)
	rootCmd.AddCommand(getBalaceCmd)
	rootCmd.AddCommand(listAddressCmd)
	rootCmd.AddCommand(printBlockChainCmd)
	rootCmd.AddCommand(printTxCmd)
	rootCmd.AddCommand(sendCmd)
}
