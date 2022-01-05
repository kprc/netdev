package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

const (
	ApiPath       = "/api"
	NetDevPath    = "netdev"
	SummaryPath   = "summary"
	FoodTowerPath = "food_tower"
	WaterPath     = "water"
	WeighPath     = "weigh"
	UniphasePath  = "uniphase"
	TriphasePath  = "triphase"
	UploadFile    = "upload"
	ProxyPath     = "proxy"
	IndexSource   = "index-source"
)

const (
	Success = 0
	Failure = 1
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func PackResult(code int, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
	}
}

func (r *Result) Bytes() []byte {
	j, _ := json.Marshal(*r)
	return j
}

func NetDevPathStr(subPath string) string {
	return path.Join(ApiPath, NetDevPath, subPath)
}

func SummaryPathStr(subPath string)  string{
	return path.Join(ApiPath,SummaryPath,subPath)
}

const (
	PigPirce = "pig-price"
	PigThirdAPI = "https://hq.sinajs.cn/list=nf_LH0"
)

func ProxyAPIPath(p string) string {
	return path.Join(ApiPath,ProxyPath,p)
}

func Proxy(writer http.ResponseWriter, request *http.Request){
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	fmt.Println(request.URL.Path)

	arrPaths:=strings.Split(request.URL.Path,"/")
	if len(arrPaths) <3 {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	switch arrPaths[3] {
	case PigPirce:
		if err := proxyPigPrice(writer);err!=nil{
			writer.WriteHeader(500)
			fmt.Fprintf(writer, "get pigprice failed")
		}
		return
	}



}

func proxyPigPrice(writer http.ResponseWriter) error {
	resp,err:=http.Get(PigThirdAPI)
	if err!=nil{
		return err
	}

	if contents, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else {
		writer.WriteHeader(200)
		writer.Write(contents)
		return nil
	}
}