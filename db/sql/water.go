package sql

import (
	"database/sql"
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertWater(db *mysqlconn.NetDevDbConn, water *msg.MsgWater) error {

	t := time.Unix(water.Timestamp,0)
	if _, err := db.Exec("Insert into water_usage (room,count,createtime ) VALUES (?,?,?)",
		water.Room,
		water.Count,
		t); err != nil {

		fmt.Println(err)
		return err
	}


	return nil
}

func SelectWater(db *mysqlconn.NetDevDbConn, beginTime, endTime int64) ([]*msg.MsgDbWater,error) {
	stmt,err:=db.Prepare("select * from water_usage where UNIX_TIMESTAMP(createtime) > ? and UNIX_TIMESTAMP(createtime) < ? ")
	if err!=nil{
		return nil,err
	}
	var rows *sql.Rows

	rows,err = stmt.Query(beginTime,endTime)
	if err!=nil{
		return nil,err
	}

	var rs []*msg.MsgDbWater

	for rows.Next(){
		w:=&msg.MsgDbWater{}
		var t time.Time
		err=rows.Scan(&w.Id,&w.Room,&w.Count,&t)
		if err!=nil{
			continue
		}
		w.Timestamp = t.Unix()
		rs = append(rs,w)
	}

	return rs,nil

}

func InsertWaterBlock(db *mysqlconn.NetDevDbConn, water *msg.MsgWater) error {
	t:=time.Unix(water.Timestamp,0)
	if _, err := db.Exec("Insert into water_usage_blockchain (room,count,createtime ) VALUES (?,?,?)",
		water.Room,
		water.Count,
		t); err != nil {
		return err
	}
	return nil
}


func UpdateWaterBlock(db *mysqlconn.NetDevDbConn, id int64, height uint64, hash string) error {
	if _, err := db.Exec("Update water_usage_blockchain set blockheight = ? , hash=? where id=?",
		height, hash, id); err != nil {
		return err
	}

	return nil
}
