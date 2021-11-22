package msg

type MsgDbWater struct {
	Id        int64   `json:"id"`
	Room      string  `json:"room"`
	Count     float64 `json:"count"`
	Timestamp int64   `json:"timestamp"`
}
