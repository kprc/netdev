package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertFoodTower(db *mysqlconn.NetDevDbConn, ft *msg.MsgFoodTower) error {
	t:=time.Unix(ft.Timestamp,0)
	if _, err := db.Exec("Insert into t_food_tower (f_room,f_weight,f_createtime ) VALUES (?,?,?)",
		ft.Room,
		ft.Weight,
		t); err != nil {
		return err
	}

	return nil
}

func InsertFoodTowerBlockChain(db *mysqlconn.NetDevDbConn, ft *msg.MsgFoodTower)  error {
	//t:=time.Unix(ft.Timestamp,0)

	return nil
}