package cmd

import (
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "发送交易",
	Long:  "<FROM> <TO> <AMOUNT> <MINER> <DATA>",
	Run:   func(cmd *cobra.Command, args []string) {},
}
