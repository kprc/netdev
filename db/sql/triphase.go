package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"strconv"
)

func InsertTriphase(db *mysqlconn.NetDevDbConn, tri *msg.MsgTriphase) error {
	if _, err := db.Exec("Insert into triphase (room,count,createtime ) VALUES (?,?,?)",
		tri.Room,
		tri.Count,
		"FROM_UNIXTIME("+strconv.FormatInt(tri.Timestamp, 10)+")"); err != nil {
		return err
	}

	return nil
}
