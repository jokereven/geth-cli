package cmd

import (
	"github.com/jokereven/geth-cli/internal"
	"github.com/spf13/cobra"
)

var from string
var to string
var amount float64
var miner string
var send_data string

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send Deal",
	Long:  "<FROM> <TO> <AMOUNT> <MINER> <DATA>",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Send(from, to, amount, miner, send_data)
	},
}

func init() {
	sendCmd.Flags().StringVarP(&from, "from", "f", "", `FROM`)
	sendCmd.Flags().StringVarP(&to, "to", "t", "", `TO`)
	sendCmd.Flags().Float64VarP(&amount, "amount", "a", amount, `AMOUNT`)
	sendCmd.Flags().StringVarP(&miner, "miner", "m", "", `MINER`)
	sendCmd.Flags().StringVarP(&send_data, "send", "s", "", `DATA`)
}
