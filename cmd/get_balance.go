package cmd

import (
	"fmt"

	"github.com/jokereven/geth-cli/internal"
	"github.com/spf13/cobra"
)

var GetBalaceAddress string

var getBalaceCmd = &cobra.Command{
	Use:   "getBalance",
	Short: "Get Balace",
	Long:  "To Get Balace",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("value: %s, type: %T", GetBalaceAddress, GetBalaceAddress)
		internal.GetBalance(GetBalaceAddress)
	},
}

func init() {
	getBalaceCmd.Flags().StringVarP(&GetBalaceAddress, "address", "a", "", `获取地址余额`)
}
