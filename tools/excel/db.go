package excel

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

const (
	mysqlUser = "root"
	mysqlPasswd = "Rickey@123"
	mysqlHost = "localhost"
	mysqlPort = 3306
	mysqlDbName = "netdev"
)


type MysqlDB struct {
	user string
	passwd string
	host string
	port int
	dbName string
	db *sql.DB
}

func stringTriAssign(a,b string) string {
	if a == ""{
		return b
	}

	return a
}

func intTriAssign(a,b int) int  {
	if a == 0{
		return b
	}
	return a
}

func NewMysqlDb(user,password,host,dbName string,port int) *MysqlDB {
	_=mysql.Config{}
	return &MysqlDB{
		user: stringTriAssign(user,mysqlUser),
		passwd: stringTriAssign(password,mysqlPasswd),
		host: stringTriAssign(host,mysqlHost),
		port: intTriAssign(port,mysqlPort),
		dbName: stringTriAssign(dbName,mysqlDbName),
	}
}


func (ei *MysqlDB)Connect() error  {
	mysqldns:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		ei.user, ei.passwd, ei.host, ei.port,ei.dbName)

	fmt.Println(mysqldns)

	db, err := sql.Open("mysql", mysqldns)
	if err != nil {
		return err
	}
	ei.db = db
	return nil
}

func (ei *MysqlDB)Close() error  {
	return ei.db.Close()
}

func (ei *MysqlDB)Insert2Electricity(earLabel string, usage float64,timeStamp int64) error  {
	sql:="Insert into electricity (iot_report_time,f_pig_code,f_electricity_usage ) VALUES (?,?,?)"
	t:=time.Unix(timeStamp,0).UTC()
	if _,err:=ei.db.Exec(sql,t,earLabel,usage);err!=nil{
		return err
	}

	return nil
}

func (ei *MysqlDB)SelectFromElectricity(earLabel string, time2 time.Time) error {
	sql:="Select count(*) from electricity where f_pig_code=? and iot_report_time=?"

	var count int
	if err:=ei.db.QueryRow(sql,earLabel,time2).Scan(&count);err!=nil{
		return err
	}

	fmt.Println(count)

	return nil
}

