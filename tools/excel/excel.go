package excel

import (
	"encoding/base64"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
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
	
	//for _, row := range rows {
	lastDay,lastMonth := 0, 8
	for idx, colCell := range rows[0] {
		if colCell != "" && base64.StdEncoding.EncodeToString([]byte(colCell)) != "IA=="{
			y,lastDay,lastMonth,tms:=CellDate(colCell,lastMonth,lastDay)
			e.ed[idx] = &ExcelDate{
				Y:y,
				M: lastMonth,
				D: lastDay,
				TMS: tms,
			}
		}
	}

	for _,row:=range rows[1:]{
		label:=row[0]
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
						//fmt.Println(label,usage,v.TMS)
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

