package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertWeigh(db *mysqlconn.NetDevDbConn, weigh *msg.MsgWeigh) error {
	t := time.Unix(weigh.Timestamp, 0).UTC()
	if _, err := db.Exec("Insert into t_weigh ("+
		"f_room,f_mao_weigh,"+
		"f_pi_weigh,f_jing_weigh,f_unit,f_createtime ) VALUES (?,?,?,?,?,?)",
		weigh.Room,
		weigh.Mao,
		weigh.Pi,
		weigh.Jing,
		weigh.Unit,
		t); err != nil {
		return err
	}

	return nil
}
