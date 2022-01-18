package mockup

import (
	"fmt"
	"github.com/kprc/netdev/config"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"github.com/kprc/netdev/server/webserver/api"
	"github.com/kprc/netdev/server/webserver/msg"
	"strings"
	"time"
)

const (
	localhost = "http://localhost:"
)

var quit chan struct{}

func getNetDevWebPort() string {
	conf := config.GetNetDevConf()

	arrPort := strings.Split(conf.WConf.ListenServer, ":")
	if len(arrPort) != 2 {
		return "0"
	}
	return arrPort[1]
}

func postPath(subPath string) string {
	return localhost + getNetDevWebPort() + api.NetDevPathStr(subPath)
}

func postSummaryPath(subpath string) string {
	return localhost + getNetDevWebPort() + api.SummaryPathStr(subpath)
}

func getLastData(db *mysqlconn.NetDevDbConn, room string) (lastWater, lastFood, lastTri, lastUni float64) {
	lastWater = GetLastWaterUsage(db, room)
	lastFood = GetLastFoodUsage(db, room)
	lastTri = GetLastTriEleUsage(db, room)
	lastUni = GetLastUniEleUsage(db, room)
	return
}

const (
	inaccuracy         = 2
	inaccuracyInterval = 100
	oneDaySecond       = 86400
	oneHourSecond      = 3600
	pigHouse           = 2
	electricUsage      = 25
	waterUsage         = 26
	foodUsage          = 27
)

func posOneDayData(lastRound *int64) error {

	t := time.Now().UTC().Unix()+ 8*3600


	//if *lastRound == 0 || (t%oneDaySecond < inaccuracy && (t-*lastRound) > inaccuracyInterval) {
	if t%oneDaySecond < inaccuracy && (t-*lastRound) > inaccuracyInterval {
		*lastRound = t
	} else {
		return nil
	}
	db := mysqlconn.NewMysqlDb()
	if err := db.Connect(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	tBegin := (t/oneDaySecond )*oneDaySecond*1000 - 8*3600000

	if houses, err := sql.SelectAllPigHouse(db); err != nil {
		fmt.Println(err.Error())
		return err
	} else {

		for _, house:=range houses{
			lastWater, lastFood, lastTri, lastUni := getLastData(db, house)
			if (lastWater + lastFood + lastTri + lastUni) == 0 {
				return nil
			}
			if err = IndexSourceInsert(tBegin, pigHouse, house, electricUsage, lastTri+lastUni); err != nil {
				fmt.Println(err)
			}
			if err = IndexSourceInsert(tBegin, pigHouse, house, waterUsage, lastWater); err != nil {
				fmt.Println(err)
			}
			if err = IndexSourceInsert(tBegin, pigHouse, house, foodUsage, lastFood); err != nil {
				fmt.Println(err)
			}

		}

	}
	return nil
}

const (
	LocalDb = 0
	RemoteDb = 1
)

func sumPigUsages(pigUsages map[string]float64, oneHousePig map[string]struct{}) float64  {
	sum:=float64(0)
	for k,u:=range pigUsages{
		if _,ok:=oneHousePig[k];ok{
			sum += u
		}
	}
	return sum
}

func postElectricity(now time.Time,dbLocal,dbRemote *mysqlconn.NetDevDbConn, pigHouses map[int]string)  {

	l,_:=time.LoadLocation("Asia/Shanghai")
	nowL := now.In(l)

	y:=nowL.Year()
	m:=nowL.Month()
	d:=nowL.Day()

	ts:=time.Date(y,m,d,0,0,0,0,time.UTC)
	for{
		allElectricitys, err:= sql.SelectElectricity(dbRemote,&ts)
		if err!=nil || len(allElectricitys) == 0{
			fmt.Println(err)
			return
		}
		tsLocal := time.Date(y,m,d,0,0,0,0,l)
		for pigIdx,house:=range pigHouses{
			var beginat int64
			if beginat,err=sql.SelectIndexData(dbLocal,house,25,tsLocal.UnixMilli());(err!=nil)|| (beginat == 0){
				var codes map[string]struct{}
				codes,err=sql.SelectSwinePigCodes(dbLocal,pigIdx)
				if err!=nil{
					fmt.Println("select codes failed",err)
					continue
				}
				usage:=sumPigUsages(allElectricitys,codes)
				if usage == 0{
					continue
				}
				//fmt.Println(usage,house)
				if err = insert2IndexSource(dbLocal,tsLocal.UnixMilli(),usage,25,2,house);err!=nil{
					fmt.Println(err)
				}
			}else {
				fmt.Println(beginat,"------")
				return
			}
		}
		timestamp := ts.Unix() - 86400
		ts = time.Unix(timestamp,0).UTC()
		y = ts.In(l).Year()
		m = ts.In(l).Month()
		d = ts.In(l).Day()
	}
}

func insert2IndexSource(db *mysqlconn.NetDevDbConn, beginAt int64, usage float64, fType int, cateType int, cateName string) error {

	t := time.Now().UTC().UnixMilli() - 8*3600000

	is := &msg.MsgIndexSource{
		Version:      0,
		BeginTime:    beginAt,
		Category:     cateType,
		CategoryCode: cateName,
		FType:        fType,
		FValue:       usage,
		State:        0,
		BaseValue:    0,
		Deleted:      0,
		CreateAt:     t,
		UpdateAt:     t,
	}

	if err:=sql.InsertIndexSource(db,is);err!=nil{
		return err
	}

	return nil

}


func postWater(now time.Time,dbLocal,dbRemote *mysqlconn.NetDevDbConn, pigHouses map[int]string)  {

	l,_:=time.LoadLocation("Asia/Shanghai")
	nowL := now.In(l)

	y:=nowL.Year()
	m:=nowL.Month()
	d:=nowL.Day()

	ts:=time.Date(y,m,d,0,0,0,0,time.UTC)
	for{

		allWaters, err:= sql.SelectWaters(dbRemote,&ts)
		if err!=nil||len(allWaters) == 0{
			fmt.Println(err)
			return
		}
		tsLocal := time.Date(y,m,d,0,0,0,0,l)
		for pigIdx,house:=range pigHouses{
			var beginat int64
			if beginat,err=sql.SelectIndexData(dbLocal,house,26,tsLocal.UnixMilli());(err!=nil)|| (beginat == 0){
				var codes map[string]struct{}
				codes,err=sql.SelectSwinePigCodes(dbLocal,pigIdx)
				if err!=nil{
					fmt.Println("select codes failed",err)
					continue
				}
				usage:=sumPigUsages(allWaters,codes)
				if usage == 0{
					continue
				}
				if err = insert2IndexSource(dbLocal,tsLocal.UnixMilli(),usage,26,2,house);err!=nil{
					fmt.Println(err)
				}
			}else {
				return
			}
		}
		timestamp := ts.Unix() - 86400
		ts = time.Unix(timestamp,0).UTC()
		y = ts.In(l).Year()
		m = ts.In(l).Month()
		d = ts.In(l).Day()
	}
}

func postFodder(now time.Time,dbLocal,dbRemote *mysqlconn.NetDevDbConn, pigHouses map[int]string)  {

	l,_:=time.LoadLocation("Asia/Shanghai")
	nowL := now.In(l)

	y:=nowL.Year()
	m:=nowL.Month()
	d:=nowL.Day()

	ts:=time.Date(y,m,d,0,0,0,0,time.UTC)
	for{

		allfodders, err:= sql.SelectFodders(dbRemote,&ts)
		if err!=nil || len(allfodders) == 0{
			fmt.Println(err)
			return
		}
		tsLocal := time.Date(y,m,d,0,0,0,0,l)
		for pigIdx,house:=range pigHouses{
			var beginat int64
			if beginat,err=sql.SelectIndexData(dbLocal,house,27,tsLocal.UnixMilli());(err!=nil)|| (beginat == 0){
				var codes map[string]struct{}
				codes,err=sql.SelectSwinePigCodes(dbLocal,pigIdx)
				if err!=nil{
					fmt.Println("select codes failed",err)
					continue
				}
				usage:=sumPigUsages(allfodders,codes)
				if usage == 0{
					continue
				}
				if err = insert2IndexSource(dbLocal,tsLocal.UnixMilli(),usage,27,2,house);err!=nil{
					fmt.Println(err)
				}
			}else {
				return
			}
		}
		timestamp := ts.Unix() - 86400
		ts = time.Unix(timestamp,0).UTC()
		y = ts.In(l).Year()
		m = ts.In(l).Month()
		d = ts.In(l).Day()
	}
}

func post2IndexSource(now time.Time) error {
	cfg:=config.GetNetDevConf()
	dbLocal := mysqlconn.NewMysqlDb1(cfg.Db[LocalDb])
	if err := dbLocal.Connect(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer dbLocal.Close()

	dbRemote := mysqlconn.NewMysqlDb1(cfg.Db[RemoteDb])
	if err := dbRemote.Connect(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer dbRemote.Close()
	//j,_:=json.MarshalIndent(cfg.Db[LocalDb]," ","\t")
	//fmt.Println(string(j))
	//j,_=json.MarshalIndent(cfg.Db[RemoteDb]," ","\t")
	//fmt.Println(string(j))

	pighouses,err:=sql.SelectAllPigHouse(dbLocal)
	if err!=nil{
		fmt.Println(err)
		return err
	}

	postElectricity(now,dbLocal,dbRemote,pighouses)
	//postWater(now,dbLocal,dbRemote,pighouses)
	//postFodder(now,dbLocal,dbRemote,pighouses)

	return nil



}


func post2IndexSource2(now time.Time) error {
	cfg:=config.GetNetDevConf()
	dbLocal := mysqlconn.NewMysqlDb1(cfg.Db[LocalDb])
	if err := dbLocal.Connect(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer dbLocal.Close()

	dbRemote := mysqlconn.NewMysqlDb1(cfg.Db[RemoteDb])
	if err := dbRemote.Connect(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer dbRemote.Close()
	//j,_:=json.MarshalIndent(cfg.Db[LocalDb]," ","\t")
	//fmt.Println(string(j))
	//j,_=json.MarshalIndent(cfg.Db[RemoteDb]," ","\t")
	//fmt.Println(string(j))

	pighouses,err:=sql.SelectAllPigHouse(dbLocal)
	if err!=nil{
		fmt.Println(err)
		return err
	}

	//postElectricity(now,dbLocal,dbRemote,pighouses)
	//postWater(now,dbLocal,dbRemote,pighouses)
	postFodder(now,dbLocal,dbRemote,pighouses)

	return nil



}


func post2IndexSource3(now time.Time) error {
	cfg:=config.GetNetDevConf()
	dbLocal := mysqlconn.NewMysqlDb1(cfg.Db[LocalDb])
	if err := dbLocal.Connect(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer dbLocal.Close()

	dbRemote := mysqlconn.NewMysqlDb1(cfg.Db[RemoteDb])
	if err := dbRemote.Connect(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer dbRemote.Close()
	//j,_:=json.MarshalIndent(cfg.Db[LocalDb]," ","\t")
	//fmt.Println(string(j))
	//j,_=json.MarshalIndent(cfg.Db[RemoteDb]," ","\t")
	//fmt.Println(string(j))

	pighouses,err:=sql.SelectAllPigHouse(dbLocal)
	if err!=nil{
		fmt.Println(err)
		return err
	}

	//postElectricity(now,dbLocal,dbRemote,pighouses)
	postWater(now,dbLocal,dbRemote,pighouses)
	//postFodder(now,dbLocal,dbRemote,pighouses)

	return nil



}

func postOneHourData() error {
	db := mysqlconn.NewMysqlDb()
	if err := db.Connect(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	if houses, err := sql.SelectAllPigHouse(db); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		for _,house:=range houses{
			lastWater, lastFood, lastTri, lastUni := getLastData(db, house)
			t := time.Now().UTC().Unix()

			if err = WaterUsage(house, t, &lastWater); err != nil {
				fmt.Println(err)
			}

			if err = FoodUsage(house, t, &lastFood); err != nil {
				fmt.Println(err)
			}

			if err = TriElectricUsage(house, t, &lastTri); err != nil {
				fmt.Println(err)
			}

			if err = UniElectricUsage(house, t, &lastUni); err != nil {
				fmt.Println(err)
			}

		}
	}

	return nil

}

func TimeOutLoop2() error {

	lastTime := int64(0)

	for  {
		select {
		case <-quit:
			return nil
		default:

		}

		now:=time.Now()

		if now.Unix() - lastTime < 300{
			time.Sleep(time.Second)
			continue
		}

		l,err:=time.LoadLocation("Asia/Shanghai")
		if err!=nil{
			return err
		}
		if now.In(l).Hour() == 0 && now.In(l).Minute() == 30 {
			lastTime = now.Unix()
			post2IndexSource(now)
		}
		time.Sleep(time.Second)
	}
}


func TimeOutLoop() error {

	lastPostTime := int64(0)
	lastRound := int64(0)

	for {
		select {
		case <-quit:
			return nil
		default:
			//nothing todo...
		}

		now := time.Now().UTC().Unix()

		if now-lastPostTime > oneHourSecond {
			postOneHourData()
			lastPostTime = now
		}
		posOneDayData(&lastRound)
		time.Sleep(time.Second)
	}
}

func StopTimeOutLoop() {
	quit <- struct{}{}
}
