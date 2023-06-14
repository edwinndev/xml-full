package models

import "encoding/xml"

type MeasurementResult struct {
	Name                Name                 `xml:"dcc:name"`
	UsedMethods         []UsedMethod         `xml:"dcc:usedMethods"`
	MeasuringEquipments []MeasuringEquipment `xml:"dcc:measuringEquipments"`
	InfluenceConditions []InfluenceCondition `xml:"dcc:influenceConditions"`
	Results             []Result             `xml:"dcc:results"`
}

type UsedMethod struct {
	XMLName xml.Name `xml:"dcc:usedMethod"`
	RefType string   `xml:"refType,attr"`
	Name    Name     `xml:"dcc:name"`
}

type MeasuringEquipment struct {
	XMLName     xml.Name    `xml:"dcc:measuringEquipment"`
	RefType     string      `xml:"refType,attr"`
	Name        Name        `xml:"dcc:name"`
	Description Description `xml:"dcc:description"`
}

type Description struct {
	Lang    string `xml:"lang,attr"`
	Content string `xml:"dcc:content"`
}

type InfluenceCondition struct {
	XMLName     xml.Name    `xml:"dcc:influenceCondition"`
	RefType     string      `xml:"refType,attr"`
	Name        Name        `xml:"dcc:name"`
	Description Description `xml:"dcc:description"`
	Data        []Quantity  `xml:"dcc:data"`
}

type Quantity struct {
	XMLName xml.Name          `xml:"dcc:quantity"`
	RefType string            `xml:"refType,attr"`
	Name    Name              `xml:"dcc:name"`
	Real    Real              `xml:"dcc:real"`
	Hybrid  []RealListXMLList `xml:"si:hybrid,omitempty"`
}

type Real struct {
	Value string `xml:"si:value"`
	Unit  string `xml:"si:unit"`
}

type Result struct {
	XMLName xml.Name `xml:"dcc:result"`
	RefType string   `xml:"refType,attr"`
	Name    Name     `xml:"dcc:name"`
	Data    Data     `xml:"dcc:data"`
}

type Data struct {
	List []Quantity `xml:"dcc:list"`
}

type RealListXMLList struct {
	XMLName      xml.Name `xml:"si:realListXMLList"`
	ValueXMLList string   `xml:"si:valueXMLList"`
	UnitXMLList  string   `xml:"si:unitXMLList"`
}
