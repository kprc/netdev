package msg

import (
	"fmt"
	"testing"
	"time"
)

func TestMsgWater(t *testing.T)  {
	w:=&MsgWater{}
	w.Room = "a.b.c"
	w.Count = 100.2
	w.Timestamp = time.Now().Unix()

	fmt.Println(w.String())
}
