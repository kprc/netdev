package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertUniphase(db *mysqlconn.NetDevDbConn, uni *msg.MsgUniphase) error {
	t := time.Unix(uni.Timestamp, 0).UTC()
	if _, err := db.Exec("Insert into t_uniphase ("+
		"f_room,f_count,f_createtime ) VALUES (?,?,?)",
		uni.Room,
		uni.Count,
		t); err != nil {
		return err
	}

	return nil
}

func SelectUniEUsage(db *mysqlconn.NetDevDbConn, room string) (float64, error) {
	sql := "select f_count from t_uniphase where f_room = ? order by f_id desc limit 1"

	count := float64(0)

	if err := db.QueryRow(sql, room).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
