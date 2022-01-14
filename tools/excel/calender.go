package excel

import (
	"strconv"
	"strings"
	"time"
)

func CellDate(cell string,lastMonth int,lastDay int) (year,day,month int, timestamp int64) {
	y:=strings.Split(cell,"年")
	m:=strings.Split(y[1],"月")
	d:=strings.Split(m[1],"日")
	month = lastMonth

	day,_ = strconv.Atoi(d[0])
	year,_ = strconv.Atoi(y[0])

	if day < lastDay{
		month ++
		if month == 13{
			month = 1
		}
	}

	l,_:=time.LoadLocation("UTC")

	t:=time.Date(year,time.Month(month),day,0,0,0,0,l)

	timestamp = t.Unix()

	return
}
