package sql

import (
	"context"
	"github.com/kprc/netdev/db/mysqlconn"
	"log"
	"time"
)

func SelectElectricity(db *mysqlconn.NetDevDbConn, tm *time.Time) (map[string]float64, error) {
	sql := "select f_pig_code,f_electricity_usage from electricity where iot_report_time = ?"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, sql)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, *tm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usages := make(map[string]float64)
	for rows.Next() {
		var ph []byte
		var usage float64
		if err := rows.Scan(&ph, &usage); err != nil {
			return nil, err
		}
		usages[string(ph)] = usage
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usages, nil

}
