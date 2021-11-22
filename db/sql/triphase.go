package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertTriphase(db *mysqlconn.NetDevDbConn, tri *msg.MsgTriphase) error {
	t:=time.Unix(tri.Timestamp,0)
	if _, err := db.Exec("Insert into triphase (room,count,createtime ) VALUES (?,?,?)",
		tri.Room,
		tri.Count,
		t); err != nil {
		return err
	}

	return nil
}
