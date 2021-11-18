package cmdservice

import (
	"github.com/kprc/nbsnetwork/tools"
	"github.com/kprc/netdev/cmd/pbs"
	"github.com/kprc/netdev/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func StartCmdService() {

	if b := tools.FileExists(config.CmdSockFile()); b {
		os.RemoveAll(config.CmdSockFile())
	}

	l, err := net.Listen("unix", config.CmdSockFile())
	if err != nil {
		panic(err)
	}

	cmdServer := grpc.NewServer()

	pbs.RegisterCmdServiceServer(cmdServer, cmdServerInstance)

	reflection.Register(cmdServer)

	if err := cmdServer.Serve(l); err != nil {
		panic(err)
	}
}
