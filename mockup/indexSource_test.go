package mockup

import (
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"github.com/kprc/netdev/server/webserver/msg"
	"testing"
	"time"
)

func TestIndexSourceInsert(t *testing.T) {

	if err := IndexSourceInsert(time.Now().UTC().UnixMilli(), 2, "34", 25, 231.00); err != nil {
		panic(err)
	}

	fmt.Println("success")
}

func TestIndexSourceInsert2(t *testing.T) {
	tm := time.Now().UTC().UnixMilli()

	is := &msg.MsgIndexSource{
		Version:      0,
		BeginTime:    time.Now().UTC().UnixMilli(),
		Category:     2,
		CategoryCode: "34",
		FType:        25,
		FValue:       231.00,
		State:        0,
		BaseValue:    0,
		Deleted:      0,
		CreateAt:     tm,
		UpdateAt:     tm,
	}

	db := mysqlconn.NewMysqlDb()

	if err := db.Connect(); err != nil {
		panic(err)
	}
	defer db.Close()
	if err := sql.InsertIndexSource(db, is); err != nil {
		panic(err)
	}
	fmt.Println("success")
}

func TestIndexSourceBeginTime(t *testing.T)  {
	db := mysqlconn.NewMysqlDb()
	if err := db.Connect(); err != nil {
		panic(err)
	}
	defer db.Close()

	if at,err:=sql.SelectLatestInsertData(db,"CQAN2001",25);err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(at)
	}

}