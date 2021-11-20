package api

import (
	"encoding/json"
	"fmt"
	"github.com/kprc/nbsnetwork/tools"
	"strings"
	"testing"
)

func TestNetDevPathStr(t *testing.T) {
	fmt.Println(NetDevPathStr(FoodTowerPath))
}

func TestSplitAddr(t *testing.T)  {
	arr:=strings.Split(":1222",":")

	fmt.Println(len(arr))

	if len(arr) > 1{
		fmt.Println(arr[1])
	}
}


type TestStruct struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func (ts *TestStruct)save() error  {
	j,_:=json.Marshal(*ts)

	return tools.Save2File(j,"./teststruct")

}

func defaultTS() *TestStruct  {
	return &TestStruct{
		A: "aaa",
		B: "bbb",
		C: "ccc",
	}
}

func (ts *TestStruct)load() error  {
	data,err:=tools.OpenAndReadAll("./teststruct")
	if err!=nil{
		return err
	}

	return json.Unmarshal(data,ts)

}

func (ts *TestStruct)String() string  {
	j,_:=json.MarshalIndent(*ts,"\t"," ")
	return string(j)
}

func TestTestStruct(t *testing.T)  {
	d:=defaultTS()

	d.save()

	fmt.Println(d.String())

}

func TestTestStructCmp(t *testing.T)  {
	d:=defaultTS()

	fmt.Println(d.String())

	if err:=d.load();err!=nil{
		fmt.Println(err)
	}

	fmt.Println(d.String())
}