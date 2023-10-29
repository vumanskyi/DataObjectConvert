package parser

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	_const "github.com/vumanskyi/data-object-convert/pkg/const"
)

var parserTypeNotFound = errors.New("parser type not found")

type Parser interface {
	Parse(t, source string) (map[string]interface{}, error)
}

type StructParser struct {
}

func NewStructParser() *StructParser {
	return &StructParser{}
}

func (p *StructParser) Parse(t, source string) (map[string]interface{}, error) {
	switch t {
	case _const.JSON_FORMAT:
		return jsonParser(source)
	case _const.XML_FORMAT:
		return xmlParser(source)
	case _const.YAML_FORMAT:
		return yamlParser(source)
	}

	return nil, parserTypeNotFound
}

func jsonParser(source string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(source), &data)
	return data, err
}

func xmlParser(source string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := xml.Unmarshal([]byte(source), &data)
	return data, err
}

func yamlParser(source string) (map[string]interface{}, error) {
	panic("In progress")
}
