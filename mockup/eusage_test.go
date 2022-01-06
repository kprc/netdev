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
