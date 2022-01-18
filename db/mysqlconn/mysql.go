package mysqlconn

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kprc/netdev/config"
)

type NetDevDbConn struct {
	user string
	passwd string
	host string
	dbName string
	driver string
	*sql.DB
}

func NewMysqlDb() *NetDevDbConn {
	dbconf:=config.GetNetDevConf().Db[0]
	return &NetDevDbConn{
		user: dbconf.User,
		passwd: dbconf.Passwd,
		host: dbconf.Host,
		dbName: dbconf.DbName,
		driver: dbconf.Driver,
	}
}

func NewMysqlDb1(conf *config.DatabaseConf) *NetDevDbConn  {
	return &NetDevDbConn{
		user: conf.User,
		passwd: conf.Passwd,
		host: conf.Host,
		dbName: conf.DbName,
		driver: conf.Driver,
	}
}

func NewMysqlDb2(user,passwd,host,dbname,driver string) *NetDevDbConn {
	return &NetDevDbConn{
		user: user,
		passwd: passwd,
		host: host,
		dbName: dbname,
		driver: driver,
	}
}

func (ndb *NetDevDbConn) Connect() error {
	mysqldns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		ndb.user, ndb.passwd, ndb.host, ndb.dbName)

	fmt.Println(mysqldns)

	db, err := sql.Open(ndb.driver, mysqldns)
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
