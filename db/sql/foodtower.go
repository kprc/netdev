package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"strconv"
)

func InsertFoodWater(db *mysqlconn.NetDevDbConn, ft *msg.MsgFoodTower) error {
	if _, err := db.Exec("Insert into food_tower (room,weight,timestamp ) VALUES (?,?,?)",
		ft.Room,
		ft.Weight,
		"FROM_UNIXTIME("+strconv.FormatInt(ft.Timestamp, 10)+")"); err != nil {
		return err
	}

	return nil
}
