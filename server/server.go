package server

import (
	"errors"
	"github.com/kprc/netdev/config"
	"github.com/kprc/netdev/server/udpserver"
	"github.com/kprc/netdev/server/webserver"
	"sync"
)

type NetDevServer struct {
	udpServer *udpserver.NetDevUdpServer
	webServer *webserver.NetDevWebServer
}


var serverInstance *NetDevServer
var serverInstanceOnceDo sync.Once

func GetServerInstance() *NetDevServer  {
	if serverInstance == nil{
		serverInstanceOnceDo.Do(func() {
			serverInstance = newServer()
		})
	}

	return serverInstance;

}

func newServer() *NetDevServer {
	cfg:=config.GetNetDevConf()

	us:=udpserver.NewNetDevUdpServer(cfg.UConf.ListenServer)

	ws:=webserver.NewNetDevWebServer(cfg.WConf.ListenServer)

	return &NetDevServer{
		udpServer: us,
		webServer: ws,
	}
}




func (ns *NetDevServer)StartDaemon() error  {
	if err:=ns.udpServer.Start();err!=nil{
		return err
	}
	if err:=ns.webServer.Start();err!=nil{
		return err
	}

	return nil
}

func (ns *NetDevServer)StopDaemon() error  {

	err1 := ns.webServer.ShutDown()

	err2 :=ns.udpServer.ShutDown()

	errMsg := ""

	if err1 != nil{
		errMsg = err1.Error()
	}

	if err2 != nil{
		if errMsg == ""{
			errMsg = err2.Error()
		}else {
			errMsg = "\r\n"+err2.Error()
		}
	}

	if errMsg != ""{
		return errors.New(errMsg)
	}
	return nil
}