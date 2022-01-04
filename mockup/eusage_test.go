package mockup

import (
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"strings"

	"testing"
	"time"
)

func getlastTriEleUsage() float64  {
	db:=mysqlconn.NewMysqlDb()

	if err:=db.Connect();err!=nil{
		panic(err)
	}
	defer db.Close()

	if last,err:=sql.SelectTriEUsage(db,"34");err!=nil {
		if b:=strings.Contains(err.Error(),"no rows in result set");b{
			return 0
		}
		panic(err)
	}else {
		return last
	}
}

func getlastUniEleUsage() float64  {
	db:=mysqlconn.NewMysqlDb()

	if err:=db.Connect();err!=nil{
		panic(err)
	}
	defer db.Close()

	if last,err:=sql.SelectUniEUsage(db,"34");err!=nil {
		if b:=strings.Contains(err.Error(),"no rows in result set");b{
			return 0
		}
		panic(err)
	}else {
		return last
	}
}


func TestTriElectricUsage(t *testing.T) {

	last := getlastTriEleUsage()

	if err:=TriElectricUsage("34",time.Now().Unix(),&last);err!=nil{
		panic(err)
	}
	fmt.Println("success")
}

func TestUniElectricUsage(t *testing.T) {
	last:=getlastUniEleUsage()
	if err:=UniElectricUsage("34",time.Now().Unix(),&last);err!=nil{
		panic(err)
	}
	fmt.Println("success")
}