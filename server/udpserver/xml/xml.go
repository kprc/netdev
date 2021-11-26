package xml

import (
	"encoding/xml"
)

type Label struct {
	XMLName xml.Name `xml:"label"`
	ID      string   `xml:"id,attr"`
	X       int      `xml:"x"`
	Y       int      `xml:"y"`
	Attr    string   `xml:"attr"`
	Extend  string   `xml:"extend"`
}

type XMLLabels struct {
	XMLName xml.Name `xml:"labels"`
	Version string   `xml:"ver,attr"`
	Map     string   `xml:"map,attr"`
	Labels  []Label  `xml:"label"`
}

func Decode(data []byte) (*XMLLabels, error) {
	v := &XMLLabels{}
	err := xml.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
