package mysqlconn

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kprc/netdev/config"
)

type NetDevDbConn struct {
	*sql.DB
}

func NewMysqlDb() *NetDevDbConn {
	return &NetDevDbConn{}
}

func (ndb *NetDevDbConn) Connect() error {
	cfg := config.GetNetDevConf()

	mysqldns := fmt.Sprintf("%s:%s@/%s?charset=utf8",
		cfg.Db.User, cfg.Db.Passwd, cfg.Db.DbName)

	db, err := sql.Open(cfg.Db.Driver, mysqldns)
	if err != nil {
		return err
	} else {
		ndb.DB = db
		return err
	}
}

func (ndb *NetDevDbConn) Close() error {
	return ndb.DB.Close()
}
