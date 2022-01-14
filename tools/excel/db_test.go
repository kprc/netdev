package excel

import (
	"fmt"
	"testing"
	"time"
)

func TestDbSelectElectricity(t *testing.T)  {
	e:=NewExcelDefault()


	t2:=time.Date(2022,2,9,0,0,0,0,time.UTC)

	if err:=e.TestExcel("DDCQDX121000803",t2);err!=nil{
		fmt.Println(err)
		return
	}


}
