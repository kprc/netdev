package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"strconv"
)

func InsertUniphase(db *mysqlconn.NetDevDbConn, uni *msg.MsgUniphase) error {
	if _, err := db.Exec("Insert into uniphase (room,count,timestamp ) VALUES (?,?,?)",
		uni.Room,
		uni.Count,
		"FROM_UNIXTIME("+strconv.FormatInt(uni.Timestamp, 10)+")"); err != nil {
		return err
	}

	return nil
}
