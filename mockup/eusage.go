package mockup

import (
	"encoding/json"
	"fmt"
	"github.com/kprc/nbsnetwork/tools/httputil"
	"github.com/kprc/netdev/server/webserver/api"
	"github.com/kprc/netdev/server/webserver/msg"
	"math/rand"
	"time"
)

const(
	triEUsageBase int = 100
	uniEUsageBase int = 90
)

func TriElectricUsage(room string, timestamp int64) error {
	rand.Seed(time.Now().UnixNano())
	count:= float64(triEUsageBase) + (float64(rand.Intn(10000)) / 100.0)

	triEU := &msg.MsgTriphase{
		MsgWater:msg.MsgWater{
			Room: room,
			Timestamp: timestamp,
			Count: count,
		},
	}

	j,_:=json.Marshal(triEU)

	hp := httputil.NewHttpPost(nil,true,2,2)
	ret,code,err:=hp.ProtectPost(postPath(api.TriphasePath),string(j))
	if err!=nil{
		fmt.Println("post tri electric usage error",err)
	}

	if code != 200{
		fmt.Println("post tri electric usage error, code is ",code)
	}

	if code == 200{
		fmt.Println("post tri electric usage success, result is ",ret)
	}

	return nil
}

func UniElectricUsage(room string, timestamp int64) error  {
	rand.Seed(time.Now().UnixNano())
	count:= float64(uniEUsageBase) + (float64(rand.Intn(9000)) / 100.0)

	uniEU := &msg.MsgUniphase{
		MsgWater:msg.MsgWater{
			Room: room,
			Timestamp: timestamp,
			Count: count,
		},
	}

	j,_:=json.Marshal(uniEU)

	hp := httputil.NewHttpPost(nil,true,2,2)
	ret,code,err:=hp.ProtectPost(postPath(api.UniphasePath),string(j))
	if err!=nil{
		fmt.Println("post uni electric usage error",err)
	}

	if code != 200{
		fmt.Println("post uni electric usage error, code is ",code)
	}

	if code == 200{
		fmt.Println("post uni electric usage success, result is ",ret)
	}

	return nil
}