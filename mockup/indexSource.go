package mockup

import (
	"encoding/json"
	"fmt"
	"github.com/kprc/nbsnetwork/tools/httputil"
	"github.com/kprc/netdev/server/webserver/api"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)


func IndexSourceInsert(beginTime int64, category int, categoryCode string,ftype int, fvalue float64) error {

	t:=time.Now().UnixMilli()

	is:=&msg.MsgIndexSource{
		Version: 0,
		BeginTime: beginTime,
		Category: category,
		CategoryCode: categoryCode,
		FType: ftype,
		FValue: fvalue,
		State: 0,
		BaseValue: 0,
		Deleted: 0,
		CreateAt: t,
		UpdateAt: t,
	}

	j,_:=json.Marshal(is)

	hp := httputil.NewHttpPost(nil,true,2,2)
	ret,code,err:=hp.ProtectPost(postSummaryPath(api.IndexSource),string(j))
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