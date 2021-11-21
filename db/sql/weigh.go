package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"strconv"
)

func InsertWeigh(db *mysqlconn.NetDevDbConn, weigh *msg.MsgWeigh) error {
	if _, err := db.Exec("Insert into weigh (room,mao_weigh,pi_weigh,jing_weigh,unit,timestamp ) VALUES (?,?,?,?,?,?)",
		weigh.Room,
		weigh.Mao,
		weigh.Pi,
		weigh.Jing,
		weigh.Unit,
		"FROM_UNIXTIME("+strconv.FormatInt(weigh.Timestamp, 10)+")"); err != nil {
		return err
	}

	return nil
}
