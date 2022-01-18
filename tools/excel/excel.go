package excel

import (
	"encoding/base64"
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
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

func (e *Excel)ReadAndInsertElectricity() error  {

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

func (e *Excel)ReadAndInsertFodder() error  {

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
					if _,err:=strconv.ParseFloat(colCell,64);err!=nil{
						fmt.Println(err)
						continue
					}else{
						//if err=e.db.Insert2Electricity(label,usage,v.TMS);err!=nil{
						//	fmt.Println(err)
						//}
						var usage float64
						var ftype int

						if ftype,err = e.db.SelectPigType(label);err!=nil || ftype <1 || ftype > 4{
							fmt.Println("get ftype from t_swine",err)
							continue
						}

						usage = randFodderUsage(ftype)

						if err = e.db.Insert2Fodder(label,usage,v.TMS);err!=nil{
							fmt.Println(err)
						}
					}

				}
			}
		}
	}

	return err
}

const(
	pig1Min = 0
	pig1Max = 55
	pig2Min = 56
	pig2Max = 137
	pig3Min = 138
	pig3Max = 246
	pig4Min = 247
	pig4Max = 335
)

var  pigFoddr = [4][2]int{
	{pig1Min,pig1Max},
	{pig2Min,pig2Max},
	{pig3Min,pig3Max},
	{pig4Min,pig4Max},
}

func randFodderUsage(ftype int) (usage float64)  {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(pigFoddr[ftype-1][1] - pigFoddr[ftype-1][0])

	n += pigFoddr[ftype-1][0]

	usage = float64(n) / 100.0

	return
}

func randWaterUsage(ftype int) (usage float64)  {
	usage = randFodderUsage(ftype)
	rand.Seed(time.Now().UnixNano())
	n:=rand.Intn(5)

	n += 15

	usage = (usage * float64(n) ) / 10.0

	return
}

func (e *Excel)ReadAndInsertWater() error  {

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
					if _,err:=strconv.ParseFloat(colCell,64);err!=nil{
						fmt.Println(err)
						continue
					}else{

						var usage float64
						var ftype int

						if ftype,err = e.db.SelectPigType(label);err!=nil || ftype <1 || ftype > 4{
							fmt.Println("get ftype from t_swine",err)
							continue
						}

						usage = randWaterUsage(ftype)

						if err = e.db.Insert2Water(label,usage,v.TMS);err!=nil{
							fmt.Println(err)
						}
						//if err=e.db.Insert2Electricity(label,usage,v.TMS);err!=nil{
						//	fmt.Println(err)
						//}
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