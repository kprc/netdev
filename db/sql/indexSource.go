package sql

import (
	"context"
	"fmt"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"log"
	"time"
)

func InsertIndexSource(db *mysqlconn.NetDevDbConn, is *msg.MsgIndexSource) error {
	//tcreate := time.Now().UTC().UnixMilli()

	_, err := db.Exec("Insert into t_index_source ("+
		"f_version,"+
		"f_begin_at,"+
		"f_category,"+
		"f_category_code,"+
		"f_type,"+
		"f_value,"+
		"f_basis_value,"+
		"f_state,"+
		"f_deleted,"+
		"f_created_at,"+
		"f_updated_at"+
		") VALUES (?,?,?,?,?,?,?,?,?,?,?)",
		is.Version,
		is.BeginTime,
		is.Category,
		is.CategoryCode,
		is.FType,
		is.FValue,
		is.BaseValue,
		is.State,
		is.Deleted,
		is.CreateAt,
		is.UpdateAt)

	if err != nil {
		return err
	}
	return nil
}

func UpdateIndexSource(db *mysqlconn.NetDevDbConn, id int64, is *msg.MsgIndexSource) error {
	tupdate := time.Now().UTC().UnixMilli()
	_, err := db.Exec("Update t_index_source set "+
		"f_state =?, "+
		"f_deleted=?, "+
		"f_updated_at=? "+
		"where f_id=?",
		is.State,
		is.Deleted,
		tupdate,
		id)

	if err != nil {
		return err
	}
	return nil
}

func SelectAllPigHouse(db *mysqlconn.NetDevDbConn) (map[int]string, error) {
	query := `SELECT f_pig_house_code,f_id FROM t_pig_house where f_pig_house_code in (select distinct(f_pig_house_code) from t_pig_house) `
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	houses := make(map[int]string)
	for rows.Next() {
		var ph []byte
		var id int
		if err := rows.Scan(&ph,&id); err != nil {
			return nil, err
		}
		houses[id] = string(ph)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return houses, nil
}

func SelectLatestInsertData(db *mysqlconn.NetDevDbConn,house string, ftype int) (int64,error){
	sql:="select f_begin_at from t_index_source where f_category = 2 and  f_type=? and f_category_code=? order by f_begin_at desc limit 1"

	row:=db.QueryRow(sql,ftype, house)

	t:=int64(0)

	if err:=row.Scan(&t);err!=nil{
		return 0,err
	}

	return t,nil
}

func SelectIndexData(db *mysqlconn.NetDevDbConn,house string, ftype int, beginAt int64) (int64, error) {
	fmt.Println(ftype,beginAt,house)
	sql:="Select f_begin_at from t_index_source where  f_category=2 and  f_category_code=? and f_type=? and f_begin_at=?"

	row:=db.QueryRow(sql,house,ftype,  beginAt)


	t:=int64(0)


	if err:=row.Scan(&t);err!=nil{
		fmt.Println(err)
		return 0,err
	}

	//fmt.Println(id,t,fc,string(code),ft)
	return t,nil
}
