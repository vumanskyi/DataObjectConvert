package builder

import (
	"errors"
	_const "github.com/vumanskyi/data-object-convert/pkg/const"
	"unicode"
)

var ImplementationNotFound = errors.New("implementation not found")

type Builder interface {
	Build(className string, data map[string]interface{}) string
}

type BuilderFactory interface {
	Create(t string) (Builder, error)
}

type ObjectBuilderFactory struct {
}

func NewObjectBuilderFactory() *ObjectBuilderFactory {
	return &ObjectBuilderFactory{}
}

func (obf *ObjectBuilderFactory) Create(t string) (Builder, error) {
	switch t {
	case _const.ENTITY_FORMAT:
		return &EntityBuilder{}, nil
	case _const.VALUE_OBJECT_FORMAT:
		return &ValueObject{}, nil
	case _const.DTO_FORMAT:
		return &DataTransferObject{}, nil
	}

	return nil, ImplementationNotFound
}

func getType(value interface{}) string {
	switch v := value.(type) {
	case string:
		return "string"
	case float64:
		if float64(int(v)) == v {
			return "int"
		}
		return "float"
	case int, int32, int64:
		return "int"
	case bool:
		return "bool"
	case map[string]interface{}:
		return "array"
	case []interface{}:
		return "array"
	default:
		return "object"
	}
}

func ucFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}
