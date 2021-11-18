package cmd

import (
	"github.com/kprc/nbsnetwork/tools"
	"github.com/kprc/netdev/config"
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var InitCmd = &cobra.Command{
	Use: "init",

	Short: "netdev init  node",

	Long: `usage description::TODO::`,

	Run: initNetDev,
}

func initNetDev(cmd *cobra.Command, args []string) {
	if b:=tools.FileExists(config.NetDevHome());!b{
		os.MkdirAll(config.NetDevHome(),0755)
	}else{
		fmt.Println("old node config in the user home dir!!!!")
		return
	}

	config.InitConfig()

	fmt.Println("initial node success!")

}