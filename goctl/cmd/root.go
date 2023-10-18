/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goctl/internal/cobrax"
	"goctl/internal/version"
	"goctl/rpc"
	"os"
	"runtime"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = cobrax.NewCommand("goctl")

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//cobra.AddTemplateFuncs(template.FuncMap{
	//	"blue":    blue,
	//	"green":   green,
	//	"rpadx":   rpadx,
	//	"rainbow": rainbow,
	//})

	rootCmd.Version = fmt.Sprintf(
		"%s %s/%s", version.BuildVersion,
		runtime.GOOS, runtime.GOARCH)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(rpc.RpcCmd)
}
