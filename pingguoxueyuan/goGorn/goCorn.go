package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

func print()  {
	fmt.Println("6666")
}



var HistoryOrderGoodsScript1 = &cobra.Command{
	Use:   "HistoryOrderGoodsScript",
	Short: "history order goods script",
	Run: func(cmd *cobra.Command, args []string) {
		print()
		gocron.Every(1).Second().Do(print)
		<-gocron.Start()
	},
}


func init() {
	rootCmd.AddCommand(HistoryOrderGoodsScript1)
	//TaobaoSourceGoodsScript.Flags().BoolVarP(&getTaobaoSourceGoodsScriptOnce, "TaobaoSourceGoods_once", "o", false, "exec TaobaoSourceGoods once")

}

var rootCmd = &cobra.Command{
	Use:   "scripts",
	Short: "scripts of our project",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("666")
}

func main() {
	Execute()
}