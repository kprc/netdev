package cmdclient

import (
	"context"
	"fmt"
	"github.com/kprc/netdev/cmd/pbs"
	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use: "show",

	Short: "show command",

	Long: `usage description::TODO::`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please choose sub command")
	},
}

var showConfigCmd = &cobra.Command{
	Use: "config",

	Short: "show configuration",

	Long: `usage description::TODO::`,

	Run: showConfig,
}

func InitShow() {
	ShowCmd.AddCommand(showConfigCmd)
}

func showConfig(cmd *cobra.Command, args []string) {
	cli, err := Dial2CmdServer()
	if err != nil {
		fmt.Println("can't connect to cmd service")
		return
	}
	var resp *pbs.CommonResponse

	resp, err = cli.ShowConfig(context.TODO(), &pbs.EmptyMessage{})
	if err != nil {
		fmt.Println("call show config failed")
		return
	}

	fmt.Println(resp.Msg)
}
