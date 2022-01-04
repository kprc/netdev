package mockup

import (
	"fmt"
	"github.com/kprc/netdev/config"
	"github.com/kprc/netdev/db/mysqlconn"
	"github.com/kprc/netdev/db/sql"
	"github.com/kprc/netdev/server/webserver/api"
	"strings"
	"time"
)
const(
	localhost="http://localhost:"
)
var quit chan struct{}

func getNetDevWebPort() string {
	conf:=config.GetNetDevConf()

	arrPort := strings.Split(conf.WConf.ListenServer,":")
	if len(arrPort) != 2{
		return "0"
	}
	return arrPort[1]
}

func postPath(subPath string) string {
	return localhost+getNetDevWebPort()+api.NetDevPathStr(subPath)
}

func postSummaryPath(subpath string) string  {
	return localhost+getNetDevWebPort()+api.SummaryPathStr(subpath)
}

func getLastData(db *mysqlconn.NetDevDbConn,room string) (lastWater,lastFood,lastTri,lastUni float64) {
	lastWater = GetLastWaterUsage(db, room)
	lastFood = GetLastFoodUsage(db,room)
	lastTri = GetLastTriEleUsage(db,room)
	lastUni = GetLastUniEleUsage(db,room)
	return
}

const(
	inaccuracy = 2
	inaccuracyInterval = 100
	oneDaySecond = 86400
	oneHourSecond = 3600
	pigHouse = 2
	electricUsage = 25
	waterUsage = 26
	foodUsage = 27

)

func postOneHourData(lastRound *int64) error {
	db:=mysqlconn.NewMysqlDb()
	if err:=db.Connect();err!=nil{
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	if houses,err:=sql.SelectAllPigHouse(db);err!=nil{
		fmt.Println(err.Error())
		return err
	}else{
		for i:=0;i<len(houses);i++{
			lastWater,lastFood,lastTri,lastUni := getLastData(db,houses[i])
			t:=time.Now().Unix()

			if err = WaterUsage(houses[i],t,&lastWater);err!=nil{
				fmt.Println(err)
			}

			if err = FoodUsage(houses[i],t,&lastFood);err!=nil{
				fmt.Println(err)
			}

			if err = TriElectricUsage(houses[i],t,&lastTri);err!=nil{
				fmt.Println(err)
			}

			if err = UniElectricUsage(houses[i],t,&lastUni);err!=nil{
				fmt.Println(err)
			}

			if t%oneDaySecond < inaccuracy && (t - *lastRound) > inaccuracyInterval {
				*lastRound = t

				tBegin := (t / oneDaySecond ) * oneDaySecond * 1000

				if err = IndexSourceInsert(tBegin,pigHouse,houses[i],electricUsage,lastTri+lastUni);err!=nil{
					fmt.Println(err)
				}
				if err = IndexSourceInsert(tBegin,pigHouse,houses[i],waterUsage,lastWater);err!=nil{
					fmt.Println(err)
				}
				if err = IndexSourceInsert(tBegin,pigHouse,houses[i],foodUsage,lastFood);err!=nil{
					fmt.Println(err)
				}
			}
		}
	}

	return nil

}

func TimeOutLoop() error {

	lastPostTime:=int64(0)
	lastRound := int64(0);

	for{
		select {
			case <-quit:
				return nil
		default:
			//nothing todo...
		}

		now:= time.Now().Unix()

		if now - lastPostTime > oneHourSecond{
			postOneHourData(&lastRound)
			lastPostTime = now
		}

		time.Sleep(time.Second)
	}
}

func StopTimeOutLoop()  {
	quit <- struct{}{}
}