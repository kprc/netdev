package cmdclient

import (
	"context"
	"github.com/kprc/netdev/cmd/pbs"
	"github.com/kprc/netdev/config"
	"google.golang.org/grpc"
	"net"
)

func UnixConnect(context.Context, string) (net.Conn, error) {
	unixAddress, _ := net.ResolveUnixAddr("unix", config.CmdSockFile())
	return net.DialUnix("unix", nil, unixAddress)
}

func Dial2CmdServer() (pbs.CmdServiceClient,error) {

	conn,err:=grpc.Dial(config.CmdSockFile(),grpc.WithInsecure(),grpc.WithContextDialer(UnixConnect))
	if err!=nil{
		return nil, err
	}

	client := pbs.NewCmdServiceClient(conn)

	return client,nil
}

func init()  {
	InitShow()
}