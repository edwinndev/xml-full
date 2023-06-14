package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"xml-full/models"
)

func main() {
	file, err := os.Create("certificate.xml")
	if err != nil {
		log.Fatalln(err)
	}

	data := getData()
	xmlEnc := xml.NewEncoder(file)
	xmlEnc.Indent("", "    ")
	e1 := xmlEnc.Encode(data)
	e2 := file.Close()
	if e1 != nil {
		fmt.Println(e1.Error())
	}
	if e2 != nil {
		fmt.Println(e2.Error())
	}
}

func getData() models.DigitalCalibrationCertificate {
	var lang = "en"
	return models.DigitalCalibrationCertificate{
		XmlnsDcc:          "https://ptb.de/dcc",
		XmlnsSi:           "https://ptb.de/si",
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: "https://ptb.de/dcc https://ptb.de/dcc/v3.1.2/dcc.xsd",
		SchemaVersion:     "3.1.2",
		AdministrativeData: models.AdministrativeData{
			DccSoftware: models.DccSoftware{
				Software: models.Software{
					Name:    models.Name{Lang: lang, Content: "GEMIMEG Tool"},
					Release: "14-06-2023",
				},
			},
			CoreData: models.CoreData{
				CountryCodeISO31661:      "DE",
				UsedLangCodeISO6391:      "en",
				MandatoryLangCodeISO6391: "en",
				UniqueIdentifier:         "SummerSchool2022_JohnSmith_07",
				Identifications: []models.Identification{
					{Issuer: "calibrationLaboratory", Value: "1234456789summer", Name: models.Name{Lang: lang, Content: "Order No"}},
				},
				ReceiptDate:          "14-06-2023",
				BeginPerformanceDate: "14-06-2023",
				EndPerformanceDate:   "14-06-2023",
				PerformanceLocation:  "laboratory",
			},
			Items: []models.Item{
				{
					Name: models.Name{Lang: lang, Content: "Arduino Uno"},
					Manufacturer: models.Manufacturer{Name: models.Name{
						Lang:    lang,
						Content: "Arduino",
					}},
					Model: "V 3.0",
					Identifications: []models.Identification{
						{
							Issuer: "manufacturer",
							Value:  "92389232",
							Name:   models.Name{Lang: lang, Content: "Serial Nro"},
						},
					},
				},
			},
			CalibrationLaboratory: models.CalibrationLaboratory{
				Contact: models.Contact{
					Name:  models.Name{Lang: lang, Content: "PTB"},
					Email: "summerschool@ptb.de",
					Location: models.Location{
						City:        "Braunschweig",
						CountryCode: "DE",
						PostCode:    "2323",
						Street:      "Bundesallee",
						StreetNo:    "100",
					},
				},
			},
			RespPersons: []models.RespPerson{
				{
					Person:     models.Person{Name: models.Name{Lang: lang, Content: "John Smith"}},
					MainSigner: true,
				},
			},
			Customer: models.Customer{
				Name:     models.Name{Lang: lang, Content: "Edwin Michel"},
				Email:    "edwin.dev@gmail.com",
				Location: models.Location{City: "Braunschweig"},
			},
			Statements: []models.Statement{
				{
					RefType:     "basic_recalibration",
					Declaration: models.Declaration{Lang: lang, Content: "Date when the calibration item is to be recalibrated"},
					Date:        "14-06-2023",
					RespAuthority: models.RespAuthority{
						Name:     models.Name{Lang: lang, Content: "John Smith"},
						Email:    "john.smith@ptb.de",
						Location: models.Location{City: "Braunschweig"},
					},
				},
			},
		},
		MeasurementResults: []models.MeasurementResult{
			{
				Name: models.Name{Lang: lang, Content: "John Smith"},
				UsedMethods: []models.UsedMethod{
					{
						RefType: "temperatureSensor",
						Name:    models.Name{Lang: lang, Content: "Calibration of temperature sensors"},
					},
				},
				MeasuringEquipments: []models.MeasuringEquipment{
					{
						RefType:     "basic_normalUsed",
						Name:        models.Name{Lang: lang, Content: "Beamex MC6-T"},
						Description: models.Description{Lang: lang, Content: "Multifunction Temperature Calibrator and Communicator"},
					},
				},
				InfluenceConditions: []models.InfluenceCondition{
					{
						RefType:     "basic_temperature",
						Name:        models.Name{Lang: lang, Content: "Ambient condition temperature"},
						Description: models.Description{Lang: lang, Content: "These values were not measured but were given"},
						Data: []models.Quantity{
							{
								RefType: "basic_temperatureMin",
								Name:    models.Name{Lang: lang, Content: "Temperature min"},
								Real:    models.Real{Value: "296", Unit: "\\kelvin"},
							}, {
								RefType: "basic_temperatureMax",
								Name:    models.Name{Lang: lang, Content: "Temperature max"},
								Real:    models.Real{Value: "299", Unit: "\\kelvin"},
							},
						},
					},
				},
				Results: []models.Result{
					{
						RefType: "measuringResult1",
						Name:    models.Name{Lang: lang, Content: "Measuring results"},
						Data: models.Data{
							List: []models.Quantity{
								{
									RefType: "basic_referenceValue",
									Name:    models.Name{Lang: lang, Content: "Reference value"},
									Real: models.Real{
										Value: "193",
										Unit:  "\\kelvin",
									},
									Hybrid: []models.RealListXMLList{
										{
											ValueXMLList: "32232",
											UnitXMLList:  "\\kelvin",
										}, {
											ValueXMLList: "29292",
											UnitXMLList:  "\\degreecelsius",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
