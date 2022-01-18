package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"log"
	"time"
)

func InsertWater(db *mysqlconn.NetDevDbConn, water *msg.MsgWater) error {

	t := time.Unix(water.Timestamp, 0).UTC()
	if _, err := db.Exec("Insert into t_water_usage ("+
		"f_room,f_count,f_createtime ) VALUES (?,?,?)",
		water.Room,
		water.Count,
		t); err != nil {

		fmt.Println(err)
		return err
	}

	return nil
}

func SelectWater(db *mysqlconn.NetDevDbConn, beginTime, endTime int64) ([]*msg.MsgDbWater, error) {
	stmt, err := db.Prepare("select * from t_water_usage where UNIX_TIMESTAMP(f_createtime) > ? and UNIX_TIMESTAMP(f_createtime) < ? ")
	if err != nil {
		return nil, err
	}
	var rows *sql.Rows

	rows, err = stmt.Query(beginTime, endTime)
	if err != nil {
		return nil, err
	}

	var rs []*msg.MsgDbWater

	for rows.Next() {
		w := &msg.MsgDbWater{}
		var t time.Time
		err = rows.Scan(&w.Id, &w.Room, &w.Count, &t)
		if err != nil {
			continue
		}
		w.Timestamp = t.Unix()
		rs = append(rs, w)
	}

	return rs, nil

}

func InsertWaterBlock(db *mysqlconn.NetDevDbConn, water *msg.MsgWater) error {
	t := time.Unix(water.Timestamp, 0).UTC()
	if _, err := db.Exec("Insert into t_water_usage_blockchain ("+
		"f_room,f_count,f_createtime ) VALUES (?,?,?)",
		water.Room,
		water.Count,
		t); err != nil {
		return err
	}
	return nil
}

func UpdateWaterBlock(db *mysqlconn.NetDevDbConn, id int64, height uint64, hash string) error {
	if _, err := db.Exec("Update t_water_usage_blockchain set "+
		"f_blockheight = ? , f_hash=? where f_id=?",
		height, hash, id); err != nil {
		return err
	}

	return nil
}

func SelectWaterUsage(db *mysqlconn.NetDevDbConn, room string) (float64, error) {
	sql := "select f_count from t_water_usage where f_room = ? order by f_id desc limit 1"

	count := float64(0)

	if err := db.QueryRow(sql, room).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}


func SelectWaters(db *mysqlconn.NetDevDbConn, tm *time.Time) (map[string]float64, error) {
	sql:="select f_pig_code,f_water_usage from water where iot_report_time = ?"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, sql)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx,*tm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usages := make(map[string]float64)
	for rows.Next() {
		var ph []byte
		var usage float64
		if err := rows.Scan(&ph,&usage); err != nil {
			return nil, err
		}
		usages[string(ph)] = usage
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usages, nil

}
