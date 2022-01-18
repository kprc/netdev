package sql

import (
	"context"
	"github.com/kprc/netdev/db/mysqlconn"
	"log"
	"time"
)

func SelectSwinePigCodes(db *mysqlconn.NetDevDbConn,houseId int) (map[string]struct{},error) {
	sql:="SELECT distinct(f_swine_code) FROM t_swine where f_pig_house_id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, sql)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx,houseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	houses := make(map[string]struct{})
	for rows.Next() {
		var ph []byte

		if err := rows.Scan(&ph); err != nil {
			return nil, err
		}
		houses[string(ph)] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return houses, nil
}
