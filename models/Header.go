package models

import "fmt"

type HeaderAttribute struct {
	Name  string
	Value string
}

func DefaultHeader() string {
	return GetHeader([]HeaderAttribute{{Name: "version", Value: "1.0"}, {Name: "encoding", Value: "UTF-8"}})
}

func GetHeader(attributes []HeaderAttribute) string {
	var headerAttributes = ""
	var size = len(attributes)
	for index, item := range attributes {
		if index < (size - 1) {
			headerAttributes += fmt.Sprintf(`%v="%v" `, item.Name, item.Value)
		} else {
			headerAttributes += fmt.Sprintf(`%v="%v"`, item.Name, item.Value)
		}
	}
	return fmt.Sprintf("<?xml %v?>\n", headerAttributes)
}
