package msg

import (
	"fmt"
	"testing"
	"time"
)

func TestMsgWater(t *testing.T) {
	w := &MsgWater{}
	w.Room = "a.b.c"
	w.Count = 100.2
	w.Timestamp = time.Now().UTC().Unix()

	fmt.Println(w.String())
}

func TestMsgFoodTower(t *testing.T) {
	ft := &MsgFoodTower{
		Room:      "f.t.e",
		Weight:    100.33,
		Timestamp: time.Now().UTC().Unix(),
	}

	fmt.Println(ft.String())
}

func TestMsgWeigh(t *testing.T) {
	w := &MsgWeigh{
		Room:      "weight.1.1",
		Mao:       100.2,
		Pi:        120.2,
		Jing:      99.1,
		Unit:      1,
		Timestamp: time.Now().UTC().Unix(),
	}

	fmt.Println(w.String())
}

func TestMsgRFID(t *testing.T) {
	rf := &MsgRFID{
		Room:      "rfid.1.1",
		LabelId:   "rfid.1able",
		X:         1,
		Y:         2,
		Attr:      1,
		Extend:    "aa",
		Timestamp: time.Now().UTC().Unix(),
	}

	fmt.Println(rf.String())

}

func TestMsgTriphase(t *testing.T) {

	tri := &MsgTriphase{MsgWater{
		Room:      "tri11.11",
		Count:     100.11,
		Timestamp: time.Now().UTC().Unix(),
	},
	}

	fmt.Println(tri.String())
}

func TestMsgUniphase(t *testing.T) {
	uni := &MsgUniphase{
		MsgWater{
			Room:      "uni11.11",
			Count:     100.11,
			Timestamp: time.Now().UTC().Unix(),
		},
	}

	fmt.Println(uni.String())
}
