package xml

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	file, err := os.Open("./example.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
		return
	}

	v,err:=Decode(data)

	if err != nil {
		panic(err)
		return
	}

	fmt.Println(v)

	for i:=0;i<len(v.Labels);i++{
		fmt.Println(v.Labels[i].X)
	}
}
