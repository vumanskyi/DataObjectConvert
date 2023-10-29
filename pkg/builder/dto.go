package builder

import (
	"fmt"
	"strings"
)

type DataTransferObject struct {
}

func (d *DataTransferObject) Build(data map[string]interface{}) string {
	var properties []string
	var builder strings.Builder

	for key, value := range data {
		dataType := getType(value)
		property := fmt.Sprintf("private %s $%s", dataType, key)
		properties = append(properties, property)
	}

	builder.WriteString(fmt.Sprintf("class %s\n", "TestDTO"))
	builder.WriteString("{\n")
	builder.WriteString("    public function __construct(\n")
	for _, prop := range properties {
		builder.WriteString("        " + prop + ",\n")
	}
	builder.WriteString("    ) {}\n")
	builder.WriteString("}")

	return builder.String()
}
