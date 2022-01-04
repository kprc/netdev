package mockup

import (
	"encoding/json"
	"fmt"
	"github.com/kprc/nbsnetwork/tools/httputil"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"github.com/kprc/netdev/server/webserver/api"
	"github.com/kprc/netdev/server/webserver/msg"
	"math/rand"
	"strings"
	"time"
)

const (
	foodUsageBase = 140
)


func GetLastFoodUsage(db *mysqlconn.NetDevDbConn,room string) float64  {
	if last,err:=sql.SelectFoodUsage(db,room);err!=nil {
		if b:=strings.Contains(err.Error(),"no rows in result set");b{
			return 0
		}
		panic(err)
	}else {
		return last
	}
}


func FoodUsage(room string, timestamp int64, last *float64) error {
	rand.Seed(time.Now().UnixNano())
	count:= float64(foodUsageBase) + (float64(rand.Intn(14000)) / 100.0)

	food:=&msg.MsgFoodTower{
		Room: room,
		Timestamp: timestamp,
		Weight: count+(*last),
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

	*last += count + (*last)

	return nil
}