package cmd

import (
	"github.com/jokereven/geth-cli/internal"
	"github.com/spf13/cobra"
)

var data []*internal.Transaction

var addBlockCmd = &cobra.Command{
	Use:   "addBlock",
	Short: "Add Block",
	Long:  "To Add Block",
	Run: func(cmd *cobra.Command, args []string) {
		internal.AddBlock(data)
	},
}

func init() {
	// TODO add data to block chain
	// addBlockCmd.Flags().Var(data, "d", `向区块链写入数据`)
}
