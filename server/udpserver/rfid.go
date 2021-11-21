package udpserver

import (
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"github.com/kprc/netdev/server/webserver/msg"
	"net"
)

func RfidUdpMsg(db *mysqlconn.NetDevDbConn, data []byte, peerAddr net.Addr) {
	_ = peerAddr

	//todo...
	//unmarshal data to msgrfid

	mr := &msg.MsgRFID{}

	if err := sql.InsertRfid(db, mr); err != nil {
		fmt.Println(err)
		fmt.Println("insert error", mr.String())
	}

}
