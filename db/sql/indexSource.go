package sql

import (
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/server/webserver/msg"
	"time"
)

func InsertIndexSource(db *mysqlconn.NetDevDbConn, is *msg.MsgIndexSource) error {
	tcreate:=time.Now().UnixMilli()

	_,err:=db.Exec("Insert into t_index_source (" +
		"f_version," +
		"f_begin_at," +
		"f_category," +
		"f_category_code" +
		"f_type" +
		"f_value" +
		"f_basis_value" +
		"f_state" +
		"f_deleted" +
		"f_created_at" +
		"f_updated_at" +
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
		tcreate,
		tcreate)

	if err!=nil{
		return err
	}
	return nil
}

func UpdateIndexSource(db *mysqlconn.NetDevDbConn,id int64, is *msg.MsgIndexSource) error  {
	tupdate:=time.Now().UnixMilli()
	_,err:=db.Exec("Update t_index_source set " +
		"f_state =?, " +
		"f_deleted=?, " +
		"f_updated_at=? " +
		"where f_id=?",
		is.State,
		is.Deleted,
		tupdate,
		id)

	if err!=nil{
		return err
	}
	return nil
}