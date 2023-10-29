package builder

import (
	"fmt"
	"strings"
)

type DataTransferObject struct {
}

func (d *DataTransferObject) Build(className string, data map[string]interface{}) string {
	var properties []string
	var builder strings.Builder

	for key, value := range data {
		dataType := getType(value)
		property := fmt.Sprintf("%s $%s", dataType, key)
		properties = append(properties, property)
	}

	propsLen := len(properties)
	builder.WriteString("<?php\n\ndeclare(strict_types=1);\n\n")

	builder.WriteString(fmt.Sprintf("class %s\n", className))
	builder.WriteString("{\n")
	for _, prop := range properties {
		builder.WriteString("    private " + prop + ";\n\n")
	}
	builder.WriteString("    public function __construct(\n")

	for idx, prop := range properties {
		if idx == propsLen-1 {
			builder.WriteString("        " + prop + "\n")
		} else {
			builder.WriteString("        " + prop + ",\n")
		}
	}
	builder.WriteString("    ) {\n")
	for key, _ := range data {
		builder.WriteString("        $this->" + key + " = " + key + ";\n")
	}
	builder.WriteString("    }\n\n")

	//@todo - add getters

	for key, value := range data {
		dataType := getType(value)
		builder.WriteString("    public function get" + ucFirst(key) + "(): " + dataType + "\n")
		builder.WriteString("    {\n")
		builder.WriteString("        return $this->" + key + ";\n")
		builder.WriteString("    }\n")
		builder.WriteString("\n")
	}
	builder.WriteString("}")

	return builder.String()
}
