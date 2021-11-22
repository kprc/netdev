package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertUniphase(db *mysqlconn.NetDevDbConn, uni *msg.MsgUniphase) error {
	t:=time.Unix(uni.Timestamp,0)
	if _, err := db.Exec("Insert into uniphase (room,count,createtime ) VALUES (?,?,?)",
		uni.Room,
		uni.Count,
		t); err != nil {
		return err
	}

	return nil
}
