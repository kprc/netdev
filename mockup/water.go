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
	waterUsageBase = 60
)

func GetLastWaterUsage(db *mysqlconn.NetDevDbConn,room string) float64  {
	if last,err:=sql.SelectWaterUsage(db,room);err!=nil {
		if b:=strings.Contains(err.Error(),"no rows in result set");b{
			return 0
		}
		panic(err)
	}else {
		return last
	}
}

func WaterUsage(room string,timestamp int64, last *float64) error {
	rand.Seed(time.Now().UnixNano())
	count:= float64(waterUsageBase) + (float64(rand.Intn(6000)) / 100.0)

	water:=&msg.MsgWater{
			Room: room,
			Timestamp: timestamp,
			Count: count+(*last),
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

	*last += count

	return nil
}