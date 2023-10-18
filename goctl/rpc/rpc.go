/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package rpc

import (
	"github.com/spf13/cobra"
	"goctl/internal/cobrax"
	"goctl/rpc/cli"
)

// rpcCmd represents the rpc command
var (
	RpcCmd = cobrax.NewCommand("rpc", cobrax.WithRunE(func(command *cobra.Command, strings []string) error {
		//return cli.RPCTemplate(true)
		return nil
	}))
	newCmd = cobrax.NewCommand("new", cobrax.WithRunE(cli.RPCNew), cobrax.WithArgs(cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)))

	//
	//newCmd = &cobra.Command{
	//	Use:   "new",
	//	Short: "",
	//	Long:  ``,
	//	Run:   cobraxcli.RPCNew,
	//}
)

func init() {
	RpcCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
