package cmd

import (
	"github.com/spf13/cobra"
)

var addBlockCmd = &cobra.Command{
	Use:   "addBlock",
	Short: "添加区块",
	Long:  "To Add Block",
	Run:   func(cmd *cobra.Command, args []string) {},
}
