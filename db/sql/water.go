package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"strconv"
)

func InsertWater(db *mysqlconn.NetDevDbConn, water *msg.MsgWater) error {
	if _, err := db.Exec("Insert into water_usage (room,count,timestamp ) VALUES (?,?,?)",
		water.Room,
		water.Count,
		"FROM_UNIXTIME("+strconv.FormatInt(water.Timestamp, 10)+")"); err != nil {
		return err
	}

	return nil
}

func InsertWaterBlock(db *mysqlconn.NetDevDbConn, water *msg.MsgWater) error {
	if _, err := db.Exec("Insert into water_usage_blockchain (room,count,timestamp ) VALUES (?,?,?)",
		water.Room,
		water.Count,
		"FROM_UNIXTIME("+strconv.FormatInt(water.Timestamp, 10)+")"); err != nil {
		return err
	}
	return nil
}

func UpdateWaterBlock(db *mysqlconn.NetDevDbConn, id int64, height uint64, hash string) error {
	if _, err := db.Exec("Update water_usage_blockchain set blockheight = ? , hash=? where id=?",
		height, hash, id); err != nil {
		return err
	}

	return nil
}
