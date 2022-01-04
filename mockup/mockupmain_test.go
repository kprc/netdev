package mockup

import (
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"testing"
)

func TestSelectAllPigHouse(t *testing.T)  {
	db:=mysqlconn.NewMysqlDb()
	if err:=db.Connect();err!=nil{
		panic(err)
	}
	defer db.Close()
	if phs,err:=sql.SelectAllPigHouse(db);err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(phs)
	}

}
