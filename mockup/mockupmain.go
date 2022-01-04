package mockup

import (
	"github.com/kprc/netdev/config"
	"github.com/kprc/netdev/server/webserver/api"
	"strings"
)

func getNetDevWebPort() string {
	conf:=config.GetNetDevConf()

	arrPort := strings.Split(conf.WConf.ListenServer,":")
	if len(arrPort) != 2{
		return "0"
	}
	return arrPort[1]
}

func postPath(subPath string) string {
	return "http://localhost:"+getNetDevWebPort()+api.NetDevPathStr(subPath)
}

