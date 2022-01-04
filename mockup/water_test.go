package mockup

import (
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"testing"
	"time"
)

func TestWaterUsage(t *testing.T) {
	db:=mysqlconn.NewMysqlDb()
	if err:=db.Connect();err!=nil{
		panic(err)
	}
	defer db.Close()

	last:=GetLastWaterUsage(db,"34")

	if err:=WaterUsage("34",time.Now().Unix(),&last);err!=nil{
		panic(err)
	}

	fmt.Println("success")
}
