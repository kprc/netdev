package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"strconv"
)

func InsertRfid(db *mysqlconn.NetDevDbConn, mr *msg.MsgRFID) error {
	if _, err := db.Exec("Insert into label_data (room,label_id,x,y,attr,extend,timestamp ) VALUES (?,?,?,?,?,?,?)",
		mr.Room,
		mr.LabelId,
		mr.X,
		mr.Y,
		mr.Attr,
		mr.Extend,
		"FROM_UNIXTIME("+strconv.FormatInt(mr.Timestamp, 10)+")"); err != nil {
		return err
	}

	return nil
}
