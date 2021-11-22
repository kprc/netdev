package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertFoodWater(db *mysqlconn.NetDevDbConn, ft *msg.MsgFoodTower) error {
	t:=time.Unix(ft.Timestamp,0)
	if _, err := db.Exec("Insert into food_tower (room,weight,createtime ) VALUES (?,?,?)",
		ft.Room,
		ft.Weight,
		t); err != nil {
		return err
	}

	return nil
}
