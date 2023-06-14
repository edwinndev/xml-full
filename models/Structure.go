package models

import "encoding/xml"

type DigitalCalibrationCertificate struct {
	XMLName            xml.Name            `xml:"dcc:digitalCalibrationCertificate"`
	XmlnsDcc           string              `xml:"xmlns:dcc,attr"`
	XmlnsSi            string              `xml:"xmlns:si,attr"`
	XmlnsXsi           string              `xml:"xmlns:xsi,attr"`
	XsiSchemaLocation  string              `xml:"xsi:schemaLocation,attr"`
	SchemaVersion      string              `xml:"schemaVersion,attr"`
	AdministrativeData AdministrativeData  `xml:"dcc:administrativeData"`
	MeasurementResults []MeasurementResult `xml:"dcc:measurementResults"`
}
