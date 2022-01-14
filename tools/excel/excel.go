package excel

import (
	"encoding/base64"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type ExcelDate struct {
	Y,M,D int
	TMS int64
}

type Excel struct {
	excelFile string
	sheet string
	ed map[int]*ExcelDate
	f *excelize.File
	db *MysqlDB
}

func NewExcel(efile, sheet string,user,passwd,host,dbName string,port int) *Excel {
	return &Excel{
		excelFile: efile,
		sheet: sheet,
		ed: make(map[int]*ExcelDate),
		db: NewMysqlDb(user,passwd,host,dbName,port),
	}
}

func NewExcelDefault()  *Excel {
	return &Excel{
		excelFile: "",
		sheet: "",
		ed: make(map[int]*ExcelDate),
		db: NewMysqlDb("","","","",0),
	}
}

func NewExcel2(efile, sheet string)  *Excel {
	return &Excel{
		excelFile: efile,
		sheet: sheet,
		ed: make(map[int]*ExcelDate),
		db: NewMysqlDb("","","","",0),
	}
}

func (e *Excel)Open() error {
	f, err := excelize.OpenFile(e.excelFile)
	if err != nil {
		println(err.Error())
		return err
	}

	e.f = f

	return nil
}

func (e *Excel)Close() error  {
	return e.f.Close()
}

func (e *Excel)ReadAndInsert() error  {

	rows, err := e.f.GetRows(e.sheet)
	if err!=nil{
		return err
	}

	if err:=e.db.Connect();err!=nil{
		return err
	}
	defer e.db.Close()

	var lastDay,lastMonth,y int = 0, 8,0
	var tms int64
	for idx, colCell := range rows[0] {
		if colCell != "" && base64.StdEncoding.EncodeToString([]byte(colCell)) != "IA=="{
			y,lastDay,lastMonth,tms=CellDate(colCell,lastMonth,lastDay)
			e.ed[idx] = &ExcelDate{
				Y:y,
				M: lastMonth,
				D: lastDay,
				TMS: tms,
			}
		}
	}
	var label string
	for _,row:=range rows[1:]{
		label=row[0]
		if label == "" || base64.StdEncoding.EncodeToString([]byte(label)) == "IA=="{
			continue
		}
		for idx, colCell :=range row{
			if idx == 0{
				continue
			}
			if colCell != "" && base64.StdEncoding.EncodeToString([]byte(colCell)) != "IA=="{
				if v,ok:=e.ed[idx];ok{
					if usage,err:=strconv.ParseFloat(colCell,64);err!=nil{
						fmt.Println(err)
						continue
					}else{
						if err=e.db.Insert2Electricity(label,usage,v.TMS);err!=nil{
							fmt.Println(err)
						}
					}

				}
			}
		}
	}

	return err
}

func (e *Excel)TestExcel(earLabel string, t time.Time) error  {

	if err:=e.db.Connect();err!=nil{
		return err
	}
	defer e.db.Close()

	if err:=e.db.SelectFromElectricity(earLabel,t);err!=nil{
		return err
	}

	return nil
}