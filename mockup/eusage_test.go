package mockup

import (
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"strings"

	"testing"
	"time"
)

func getlastTriEleUsage() float64 {
	db := mysqlconn.NewMysqlDb()

	if err := db.Connect(); err != nil {
		panic(err)
	}
	defer db.Close()

	if last, err := sql.SelectTriEUsage(db, "34"); err != nil {
		if b := strings.Contains(err.Error(), "no rows in result set"); b {
			return 0
		}
		panic(err)
	} else {
		return last
	}
}

func getlastUniEleUsage() float64 {
	db := mysqlconn.NewMysqlDb()

	if err := db.Connect(); err != nil {
		panic(err)
	}
	defer db.Close()

	if last, err := sql.SelectUniEUsage(db, "34"); err != nil {
		if b := strings.Contains(err.Error(), "no rows in result set"); b {
			return 0
		}
		panic(err)
	} else {
		return last
	}
}

func TestTriElectricUsage(t *testing.T) {

	last := getlastTriEleUsage()

	if err := TriElectricUsage("34", time.Now().UTC().Unix(), &last); err != nil {
		panic(err)
	}
	fmt.Println("success")
}

func TestUniElectricUsage(t *testing.T) {
	last := getlastUniEleUsage()
	if err := UniElectricUsage("34", time.Now().UTC().Unix(), &last); err != nil {
		panic(err)
	}
	fmt.Println("success")
}

func TestTimeLocal(t *testing.T) {
	fmt.Println(time.Now())
	//l,_:=time.LoadLocation("Local")
	//fmt.Println(l.String())
	fmt.Println(time.Now().Local())
}

func TestTimeTest(t *testing.T) {
	now := time.Now().UTC().Unix()

	fmt.Println(now)

	fmt.Println((now + 8*3600) % 86400)

	fmt.Println((now - 8*3600) % 86400)
	//tBegin := (t/oneDaySecond +1 )*oneDaySecond*1000 - 8*3600000
	fmt.Println((now/86400)*86400 *1000 - (8 *3600000))

	tt:=1641398418972 +  8*3600000

	fmt.Println((tt/1000/86400)*86400 )


}


func TestSelectElectricity(t *testing.T) {
	tm:=time.Date(2021,9,9,0,0,0,0,time.UTC)

	db := mysqlconn.NewMysqlDb()

	if err := db.Connect(); err != nil {
		panic(err)
	}
	defer db.Close()

	if us,err:=sql.SelectElectricity(db,&tm);err!=nil{
		fmt.Println(err)
	}else{
		for k,v:=range us{
			fmt.Println(k,v,tm.String())
		}
	}
}

func TestTimeZone(t *testing.T)  {
	tm:=time.Now()

	l1,_:=time.LoadLocation("")
	l2,_:=time.LoadLocation("Asia/Shanghai")
	l3,_:=time.LoadLocation("America/Los_Angeles")

	fmt.Println(tm.Day(),tm.Hour())
	fmt.Println(tm.In(l1).Day(),tm.In(l1).Hour())
	fmt.Println(tm.In(l2).Day(),tm.In(l2).Hour())
	fmt.Println(tm.In(l3).Day(),tm.In(l3).Hour())
	//
	//
	//tm1:=time.Unix(1640188800,0)
	//
	//fmt.Println(tm1.Year(),tm1.Month(),tm1.Day(),tm1.Hour(),tm1.Minute())
	//
	//
	//fmt.Println("====")
	//
	//now:=time.Now()
	//
	//fmt.Println(now.Unix())
	//
	//l,_:=time.LoadLocation("Asia/Shanghai")
	//nowL := now.In(l)
	//
	//y:=nowL.Year()
	//m:=nowL.Month()
	//d:=nowL.Day()
	//
	//ts:=time.Date(y,m,d,0,0,0,0,l)
	//
	//ts1:=time.Date(y,m,d,0,0,0,0,time.UTC)
	//
	//fmt.Println(ts.Unix(),ts1.Unix())

	//tm:=time.Unix(1640275200,0).UTC()
	//fmt.Println(tm.Day(),tm.Hour())
	//
	//
	//l2,_:=time.LoadLocation("Asia/Shanghai")
	//
	//tmlocal:=tm.In(l2)
	//
	//fmt.Println(tmlocal.Day(),tmlocal.Hour())

	testtm:=time.Date(2021,9,9,0,0,0,0,l2)

	fmt.Println(testtm.Unix())

}

func TestGetEletricitys(t *testing.T)  {
	//l,_:=time.LoadLocation("Asia/Shanghai")
	//nowL := now.In(l)
	//
	//y:=nowL.Year()
	//m:=nowL.Month()
	//d:=nowL.Day()

	ts:=time.Date(2021,9,9,0,0,0,0,time.UTC)

	db := mysqlconn.NewMysqlDb()

	if err := db.Connect(); err != nil {
		panic(err)
	}
	defer db.Close()

	if es,err:=sql.SelectElectricity(db,&ts);err!=nil{
		fmt.Println(err)
	}else{
		for k,v :=range es{
			fmt.Println(k,v)
		}
	}
}