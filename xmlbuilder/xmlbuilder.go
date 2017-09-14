package xmlbuilder

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

func ToXmlFromJson(jsonContent string) (string, error) {

	var jsonObject map[string]interface{}

	decoder := json.NewDecoder(strings.NewReader(jsonContent))

	if err := decoder.Decode(&jsonObject); err != nil {
		return "", err
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	encoder := xml.NewEncoder(writer)

	if err := writeXML(encoder, jsonObject); err != nil {
		return "", err
	}
	if err := writer.Flush(); err != nil {
		return "", err
	}
	return b.String(), nil

}
func writeXML(encoder *xml.Encoder, isonObject map[string]interface{}) error {
	for k, v := range isonObject {
		if err := writeNo(encoder, k, v); err != nil {
			return err
		}
	}
	return nil
}
func writeNo(encoder *xml.Encoder, tagName string, value interface{}) error {
	if val := reflect.ValueOf(value); val.Kind() == reflect.Map {
		if err := encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: tagName}}); err != nil {
			return err
		}
		for _, k := range val.MapKeys() {
			if err := writeNo(encoder, fmt.Sprintf("%v", k), val.MapIndex(reflect.ValueOf(fmt.Sprint(k))).Interface()); err != nil {
				return err
			}

		}
		if err := encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: tagName}}.End()); err != nil {
			return err
		}
		return nil
	} else {
		if err := encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: tagName}}); err != nil {
			return err
		}
		if err := encoder.EncodeToken(xml.CharData(fmt.Sprintf("%v", value))); err != nil {
			return err
		}
		if err := encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: tagName}}.End()); err != nil {
			return err
		}
		return nil
	}
}
