package udpserver

import (
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"github.com/kprc/netdev/server/udpserver/xml"
	"github.com/kprc/netdev/server/webserver/msg"
	"net"
	"time"
)

func RfidUdpMsg(db *mysqlconn.NetDevDbConn, data []byte, peerAddr net.Addr) {
	_ = peerAddr
	//unmarshal data to lables
	var (
		labels *xml.XMLLabels
		err    error
	)

	if labels, err = xml.Decode(data); err != nil {
		fmt.Println("Decode xml failed, ")
		fmt.Println("xml content:", string(data))
		return
	}

	for i := 0; i < len(labels.Labels); i++ {
		l := labels.Labels[i]
		mr := &msg.MsgRFID{}
		mr.LabelId = l.ID
		mr.Timestamp = time.Now().Unix()
		mr.X = int32(l.X)
		mr.Y = int32(l.Y)
		mr.Extend = l.Extend
		mr.Attr = int(l.Attr[0])
		if err := sql.InsertRfid(db, mr); err != nil {
			fmt.Println(err)
			fmt.Println("insert error", mr.String())
		}
	}

}
