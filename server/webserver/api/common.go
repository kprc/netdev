package api

import (
	"encoding/json"
	"path"
)

const(
	ApiPath = "/api"
	NetDevPath = "netdev"
	FoodTowerPath = "food_tower"
	WaterPath = "water"
	WeighPath = "weigh"
	UniphasePath = "uniphase"
	TriphasePath = "triphase"
)


const(
	Success = 0
	Failure = 1
)


type Result struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
}

func PackResult(code int, msg string) *Result {
	return &Result{
		Code: code,
		Msg: msg,
	}
}

func (r *Result)Bytes() []byte  {
	j,_:=json.Marshal(*r)
	return j
}

func NetDevPathStr(subPath string) string {
	return path.Join(ApiPath,NetDevPath,subPath)
}