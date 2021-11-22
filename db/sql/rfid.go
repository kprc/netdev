package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertRfid(db *mysqlconn.NetDevDbConn, mr *msg.MsgRFID) error {
	t:=time.Unix(mr.Timestamp,0)
	if _, err := db.Exec("Insert into label_data (room,label_id,x,y,attr,extend,createtime ) VALUES (?,?,?,?,?,?,?)",
		mr.Room,
		mr.LabelId,
		mr.X,
		mr.Y,
		mr.Attr,
		mr.Extend,
		t); err != nil {
		return err
	}

	return nil
}
