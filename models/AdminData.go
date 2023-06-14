package models

import (
	"encoding/xml"
)

type AdministrativeData struct {
	DccSoftware           DccSoftware           `xml:"dcc:dccSoftware"`
	CoreData              CoreData              `xml:"dcc:coreData"`
	Items                 []Item                `xml:"dcc:items"`
	CalibrationLaboratory CalibrationLaboratory `xml:"dcc:calibrationLaboratory"`
	RespPersons           []RespPerson          `xml:"dcc:respPersons"`
	Customer              Customer              `xml:"dcc:customer"`
	Statements            []Statement           `xml:"dcc:statements"`
}

type DccSoftware struct {
	Software Software `xml:"dcc:software"`
}

type Software struct {
	Name    Name   `xml:"dcc:name"`
	Release string `xml:"dcc:release"`
}

type Name struct {
	Lang    string `xml:"lang,attr"`
	Content string `xml:"dcc:content"`
}

type CoreData struct {
	CountryCodeISO31661      string           `xml:"dcc:countryCodeISO3166_1"`
	UsedLangCodeISO6391      string           `xml:"dcc:usedLangCodeISO639_1"`
	MandatoryLangCodeISO6391 string           `xml:"dcc:mandatoryLangCodeISO639_1"`
	UniqueIdentifier         string           `xml:"dcc:uniqueIdentifier"`
	Identifications          []Identification `xml:"dcc:identifications"`
	ReceiptDate              string           `xml:"dcc:receiptDate"`
	BeginPerformanceDate     string           `xml:"dcc:beginPerformanceDate"`
	EndPerformanceDate       string           `xml:"dcc:endPerformanceDate"`
	PerformanceLocation      string           `xml:"dcc:performanceLocation"`
}

type Item struct {
	Name            Name             `xml:"dcc:name"`
	Manufacturer    Manufacturer     `xml:"dcc:manufacturer"`
	Model           string           `xml:"dcc:model"`
	Identifications []Identification `xml:"dcc:identifications"`
}

type Identification struct {
	Issuer string `xml:"dcc:issuer"`
	Value  string `xml:"dcc:value"`
	Name   Name   `xml:"dcc:name"`
}

type Manufacturer struct {
	Name Name `xml:"dcc:name"`
}

type CalibrationLaboratory struct {
	Contact Contact `xml:"dcc:contact"`
}

type Contact struct {
	Name     Name     `xml:"dcc:name"`
	Email    string   `xml:"dcc:eMail"`
	Location Location `xml:"dcc:location"`
}

type Location struct {
	City        string `xml:"dcc:city"`
	CountryCode string `xml:"dcc:countryCode,omitempty"`
	PostCode    string `xml:"dcc:postCode,omitempty"`
	Street      string `xml:"dcc:street,omitempty"`
	StreetNo    string `xml:"dcc:streetNo,omitempty"`
}

type RespPerson struct {
	XMLName    xml.Name `xml:"dcc:respPerson"`
	Person     Person   `xml:"dcc:person"`
	MainSigner bool     `xml:"dcc:mainSigner"`
}

type Person struct {
	Name Name `xml:"dcc:name"`
}

type Customer struct {
	Name     Name     `xml:"dcc:name"`
	Email    string   `xml:"dcc:eMail"`
	Location Location `xml:"dcc:location"`
}

type Statement struct {
	RefType       string        `xml:"refType,attr"`
	Declaration   Declaration   `xml:"dcc:declaration"`
	Date          string        `xml:"dcc:date"`
	RespAuthority RespAuthority `xml:"dcc:respAuthority"`
}

type Declaration struct {
	Lang    string `xml:"lang,attr"`
	Content string `xml:"dcc:content"`
}

type RespAuthority struct {
	Name     Name     `xml:"dcc:name"`
	Email    string   `xml:"dcc:eMail"`
	Location Location `xml:"dcc:location"`
}
