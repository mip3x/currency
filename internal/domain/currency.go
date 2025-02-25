package domain

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type ValCurs struct {
	XMLName xml.Name       `xml:"ValCurs"`
	Valutes []SimpleValute `xml:"Valute"`
}

type SimpleValute struct {
	CharCode string
	Value    float64
}

func (sv *SimpleValute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var aux struct {
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	sv.CharCode = aux.CharCode
	valueStr := strings.Replace(aux.Value, ",", ".", 1)

	v, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return err
	}
	sv.Value = v

	return nil
}
