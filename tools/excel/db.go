package excel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/kprc/netdev/server/webserver/msg"
	"log"
	"time"
)

const (
	mysqlUser   = "root"
	mysqlPasswd = "Rickey@123"
	mysqlHost   = "localhost"
	mysqlPort   = 3306
	mysqlDbName = "netdev"
)

type MysqlDB struct {
	user   string
	passwd string
	host   string
	port   int
	dbName string
	db     *sql.DB
}

func stringTriAssign(a, b string) string {
	if a == "" {
		return b
	}

	return a
}

func intTriAssign(a, b int) int {
	if a == 0 {
		return b
	}
	return a
}

func NewMysqlDb(user, password, host, dbName string, port int) *MysqlDB {
	_ = mysql.Config{}
	return &MysqlDB{
		user:   stringTriAssign(user, mysqlUser),
		passwd: stringTriAssign(password, mysqlPasswd),
		host:   stringTriAssign(host, mysqlHost),
		port:   intTriAssign(port, mysqlPort),
		dbName: stringTriAssign(dbName, mysqlDbName),
	}
}

func (ei *MysqlDB) Connect() error {
	mysqldns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		ei.user, ei.passwd, ei.host, ei.port, ei.dbName)

	fmt.Println(mysqldns)

	db, err := sql.Open("mysql", mysqldns)
	if err != nil {
		return err
	}
	ei.db = db
	return nil
}

func (ei *MysqlDB) Close() error {
	return ei.db.Close()
}

func (ei *MysqlDB) Insert2Electricity(earLabel string, usage float64, timeStamp int64) error {
	sql := "Insert into electricity (iot_report_time,f_pig_code,f_electricity_usage ) VALUES (?,?,?)"
	t := time.Unix(timeStamp, 0).UTC()
	if _, err := ei.db.Exec(sql, t, earLabel, usage); err != nil {
		return err
	}

	return nil
}

func (ei *MysqlDB) Insert2Fodder(earLabel string, usage float64, timeStamp int64) error {
	sql := "Insert into fodder (iot_report_time,f_pig_code,f_fodder_usage ) VALUES (?,?,?)"
	t := time.Unix(timeStamp, 0).UTC()
	if _, err := ei.db.Exec(sql, t, earLabel, usage); err != nil {
		return err
	}

	return nil
}

func (ei *MysqlDB) Insert2Water(earLabel string, usage float64, timeStamp int64) error {
	sql := "Insert into water (iot_report_time,f_pig_code,f_water_usage ) VALUES (?,?,?)"
	t := time.Unix(timeStamp, 0).UTC()
	if _, err := ei.db.Exec(sql, t, earLabel, usage); err != nil {
		return err
	}

	return nil
}

func (ei *MysqlDB) SelectPigType(pigCode string) (int, error) {
	sql := "select f_type from t_swine where f_swine_code = ?"
	var ftype int
	if err := ei.db.QueryRow(sql, pigCode).Scan(&ftype); err != nil {
		return 0, err
	}

	return ftype, nil
}

func (ei *MysqlDB) SelectFromElectricity(earLabel string, time2 time.Time) error {
	sql := "Select count(*) from electricity where f_pig_code=? and iot_report_time=?"

	var count int
	if err := ei.db.QueryRow(sql, earLabel, time2).Scan(&count); err != nil {
		return err
	}

	//fmt.Println(count)

	return nil
}

func (ei *MysqlDB) SelectElectricityOnePigHouse(reportTime time.Time, house int) (float64, error) {
	sql := "select sum(f_electricity_usage) from electricity where iot_report_time = ? and" +
		" f_pig_code in (select f_swine_code from t_swine where f_pig_house_id=?)"

	var sum float64
	if err := ei.db.QueryRow(sql, reportTime, house).Scan(&sum); err != nil {
		return 0, err
	}

	return sum, nil
}

func (ei *MysqlDB) SelectAllHouse() (map[int]string, error) {
	query := `SELECT distinct(f_pig_house_code),f_id FROM t_pig_house`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := ei.db.PrepareContext(ctx, query)
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
		if err := rows.Scan(&ph, &id); err != nil {
			return nil, err
		}
		houses[id] = string(ph)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return houses, nil

}

func (ei *MysqlDB) SelectBeforeTime(category_code string, typ int) ([]time.Time, error) {
	//sql:="select min(f_begin_at) from t_index_source where f_category = 2 and f_category_code = ? and f_type = ? and f_begin_at > ? "
	return nil, nil
}

func (ei *MysqlDB) InsertIndexSource(is *msg.MsgIndexSource) error {
	_, err := ei.db.Exec("Insert into t_index_source ("+
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
