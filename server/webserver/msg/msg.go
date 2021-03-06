package msg

import "encoding/json"

type MsgFoodTower struct {
	Room      string  `json:"room"`
	Weight    float64 `json:"weight"`
	Timestamp int64   `json:"timestamp"`
}

func (ft *MsgFoodTower) Bytes() []byte {
	j, _ := json.Marshal(*ft)

	return j
}

func (ft *MsgFoodTower) String() string {
	return string(ft.Bytes())
}

type MsgWater struct {
	Room      string  `json:"room"`
	Count     float64 `json:"count"`
	Timestamp int64   `json:"timestamp"`
}

func (mw *MsgWater) Bytes() []byte {
	j, _ := json.Marshal(*mw)

	return j
}

func (mw *MsgWater) String() string {
	return string(mw.Bytes())
}

type MsgWeigh struct {
	Room      string  `json:"room"`
	Mao       float64 `json:"mao"`
	Pi        float64 `json:"pi"`
	Jing      float64 `json:"jing"`
	Unit      int     `json:"unit"`
	Timestamp int64   `json:"timestamp"`
}

func (mw *MsgWeigh) Bytes() []byte {
	j, _ := json.Marshal(*mw)

	return j
}

func (mw *MsgWeigh) String() string {
	return string(mw.Bytes())
}

type MsgUniphase struct {
	MsgWater
}
type MsgTriphase struct {
	MsgWater
}

type MsgRFID struct {
	Room      string `json:"room"`
	LabelId   string `json:"label_id"`
	X         int32  `json:"x"`
	Y         int32  `json:"y"`
	Attr      int    `json:"attr"`
	Extend    string `json:"extend"`
	Timestamp int64  `json:"timestamp"`
}

func (mr *MsgRFID) Bytes() []byte {
	j, _ := json.Marshal(*mr)

	return j
}

func (mr *MsgRFID) String() string {
	return string(mr.Bytes())
}

type MsgIndexSource struct {
	Version      int     `json:"version"`
	BeginTime    int64   `json:"begin_time"`
	Category     int     `json:"category"`
	CategoryCode string  `json:"category_code"`
	FType        int     `json:"f_type"`
	FValue       float64 `json:"f_value"`
	BaseValue    float64 `json:"base_value"`
	State        int     `json:"state"`
	Deleted      int     `json:"deleted"`
	CreateAt     int64   `json:"create_at"`
	UpdateAt     int64   `json:"update_at"`
}

func (is *MsgIndexSource) Bytes() []byte {
	j, _ := json.Marshal(*is)

	return j
}

func (is *MsgIndexSource) String() string {
	return string(is.Bytes())
}
