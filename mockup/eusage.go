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

const(
	triEUsageBase int = 100
	uniEUsageBase int = 90
)

func GetLastTriEleUsage(db *mysqlconn.NetDevDbConn,room string) float64  {
	if last,err:=sql.SelectTriEUsage(db,room);err!=nil {
		if b:=strings.Contains(err.Error(),"no rows in result set");b{
			return 0
		}
		panic(err)
	}else {
		return last
	}
}

func GetLastUniEleUsage(db *mysqlconn.NetDevDbConn, room string) float64  {
	if last,err:=sql.SelectUniEUsage(db,room);err!=nil {
		if b:=strings.Contains(err.Error(),"no rows in result set");b{
			return 0
		}
		panic(err)
	}else {
		return last
	}
}

func TriElectricUsage(room string, timestamp int64, lastCount *float64) error {
	rand.Seed(time.Now().UnixNano())
	count:= float64(triEUsageBase) + (float64(rand.Intn(10000)) / 100.0)

	triEU := &msg.MsgTriphase{
		MsgWater:msg.MsgWater{
			Room: room,
			Timestamp: timestamp,
			Count: count+(*lastCount),
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

	*lastCount = *lastCount + count

	return nil
}

func UniElectricUsage(room string, timestamp int64,lastCount *float64) error  {
	rand.Seed(time.Now().UnixNano())
	count:= float64(uniEUsageBase) + (float64(rand.Intn(9000)) / 100.0)

	uniEU := &msg.MsgUniphase{
		MsgWater:msg.MsgWater{
			Room: room,
			Timestamp: timestamp,
			Count: count+(*lastCount),
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

	*lastCount = *lastCount + count

	return nil
}

