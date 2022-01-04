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
	foodUsageBase = 140
)

func FoodUsage(room string, timestamp int64) error {
	rand.Seed(time.Now().UnixNano())
	count:= float64(foodUsageBase) + (float64(rand.Intn(14000)) / 100.0)

	food:=&msg.MsgFoodTower{
		Room: room,
		Timestamp: timestamp,
		Weight: count,
	}

	j,_:=json.Marshal(food)

	hp := httputil.NewHttpPost(nil,true,2,2)
	ret,code,err:=hp.ProtectPost(postPath(api.FoodTowerPath),string(j))
	if err!=nil{
		fmt.Println("post food usage error",err)
	}

	if code != 200{
		fmt.Println("post food usage error, code is ",code)
	}

	if code == 200{
		fmt.Println("post food usage success, result is ",ret)
	}

	return nil
}