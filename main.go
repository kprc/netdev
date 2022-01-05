package main

import (
	"fmt"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/kprc/netdev/cmd"
	"github.com/kprc/netdev/cmd/cmdclient"
	"github.com/kprc/netdev/cmd/cmdservice"
	"github.com/kprc/netdev/config"
	"github.com/kprc/netdev/mockup"
	"github.com/kprc/netdev/server"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
	"time"
)

const (
	PidFileName = ".pid"
	Version     = "0.0.1"
)

var param = struct {
	version bool
	passwd  string
}{}

var rootCmd = &cobra.Command{
	Use: "netdev",

	Short: "netdev node for service node",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

func init() {

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "netdev -v")

	rootCmd.AddCommand(cmdclient.ShowCmd)
	rootCmd.AddCommand(cmd.InitCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func mainRun(_ *cobra.Command, _ []string) {

	if b := tools.FileExists(config.NetDevHome()); !b {
		fmt.Println("please initial netdev server first!")
		return
	}

	if param.version {
		fmt.Println(Version)
		return
	}

	go server.GetServerInstance().StartDaemon()

	go cmdservice.StartCmdService()

	time.Sleep(time.Second*2)

	go mockup.TimeOutLoop()

	waitShutdownSignal()
}

func waitShutdownSignal() {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>netdev start at pid(%s)<<<<<<<<<<\n", pid)

	pidfile := path.Join(config.NetDevHome(), PidFileName)

	if err := ioutil.WriteFile(pidfile, []byte(pid), 0644); err != nil {
		fmt.Println("failed to write running pid", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2)

	sig := <-sigCh

	mockup.StopTimeOutLoop()

	server.GetServerInstance().StopDaemon()


	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
