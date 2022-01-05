package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertFoodTower(db *mysqlconn.NetDevDbConn, ft *msg.MsgFoodTower) error {
	t := time.Unix(ft.Timestamp, 0).UTC()
	if _, err := db.Exec("Insert into t_food_tower (f_room,f_weight,f_createtime ) VALUES (?,?,?)",
		ft.Room,
		ft.Weight,
		t); err != nil {
		return err
	}

	return nil
}

func InsertFoodTowerBlockChain(db *mysqlconn.NetDevDbConn, ft *msg.MsgFoodTower) error {
	//t:=time.Unix(ft.Timestamp,0)

	return nil
}

func SelectFoodUsage(db *mysqlconn.NetDevDbConn, room string) (float64, error) {
	sql := "select f_weight from t_food_tower where f_room = ? order by f_id desc limit 1"

	count := float64(0)

	if err := db.QueryRow(sql, room).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
