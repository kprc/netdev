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

const (
	waterUsageBase = 60
)

func WaterUsage(room string,timestamp int64) error {
	rand.Seed(time.Now().UnixNano())
	count:= float64(waterUsageBase) + (float64(rand.Intn(6000)) / 100.0)

	water:=&msg.MsgWater{
			Room: room,
			Timestamp: timestamp,
			Count: count,
		}

	j,_:=json.Marshal(water)

	hp := httputil.NewHttpPost(nil,true,2,2)
	ret,code,err:=hp.ProtectPost(postPath(api.WaterPath),string(j))
	if err!=nil{
		fmt.Println("post water usage error",err)
	}

	if code != 200{
		fmt.Println("post water usage error, code is ",code)
	}

	if code == 200{
		fmt.Println("post water usage success, result is ",ret)
	}

	return nil
}