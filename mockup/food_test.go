package mockup

import (
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"testing"
	"time"
)

func TestFoodUsage(t *testing.T) {
	db := mysqlconn.NewMysqlDb()
	if err := db.Connect(); err != nil {
		panic(err)
	}
	defer db.Close()

	last := GetLastFoodUsage(db, "34")

	if err := FoodUsage("34", time.Now().UTC().Unix(), &last); err != nil {
		panic(err)
	}

	fmt.Println("success")
}
