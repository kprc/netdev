package api

import (
	"fmt"
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
