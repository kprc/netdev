package cmdservice

import (
	"context"
	"encoding/json"
	"github.com/kprc/netdev/cmd/pbs"
	"github.com/kprc/netdev/config"
)

func (cs *CmdSrv)ShowConfig(context.Context, *pbs.EmptyMessage) (*pbs.CommonResponse, error)  {

	nc:=config.GetNetDevConf()

	j,_:=json.MarshalIndent(*nc,"\t"," ")

	return &pbs.CommonResponse{
		Msg: string(j),
	},nil
}