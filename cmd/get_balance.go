package cmd

import (
	"github.com/spf13/cobra"
)

var getBalaceCmd = &cobra.Command{
	Use:   "getBalance",
	Short: "获取余额",
	Long:  "To Get Balace",
	Run:   func(cmd *cobra.Command, args []string) {},
}
