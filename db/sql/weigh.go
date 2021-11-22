package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertWeigh(db *mysqlconn.NetDevDbConn, weigh *msg.MsgWeigh) error {
	t:=time.Unix(weigh.Timestamp,0)
	if _, err := db.Exec("Insert into weigh (room,mao_weigh,pi_weigh,jing_weigh,unit,createtime ) VALUES (?,?,?,?,?,?)",
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
