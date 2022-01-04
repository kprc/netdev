package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertTriphase(db *mysqlconn.NetDevDbConn, tri *msg.MsgTriphase) error {
	t:=time.Unix(tri.Timestamp,0)
	if _, err := db.Exec("Insert into t_triphase (" +
		"f_room,f_count,f_createtime ) VALUES (?,?,?)",
		tri.Room,
		tri.Count,
		t); err != nil {
		return err
	}

	return nil
}

func SelectTriEUsage(db *mysqlconn.NetDevDbConn,room string) (float64,error) {
	sql:="select f_count from t_triphase where f_room = ? order by f_id desc limit 1"

	count:=float64(0)

	if err:=db.QueryRow(sql,room).Scan(&count);err!=nil{
		return 0,err
	}

	return count,nil
}

