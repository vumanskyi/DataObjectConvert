package builder

import (
	"fmt"
	"strings"
)

type EntityBuilder struct {
}

func (e *EntityBuilder) Build(className string, data map[string]interface{}) string {
	var properties []string
	var builder strings.Builder

	for key, value := range data {
		dataType := getType(value)
		property := fmt.Sprintf("%s $%s", dataType, key)
		properties = append(properties, property)
	}

	builder.WriteString("<?php\n\ndeclare(strict_types=1);\n\n")

	builder.WriteString(fmt.Sprintf("class %s\n", className))
	builder.WriteString("{\n")
	for _, prop := range properties {
		builder.WriteString("    private " + prop + ";\n\n")
	}

	for key, value := range data {
		dataType := getType(value)
		ucFirstName := ucFirst(key)
		builder.WriteString("    public function set" + ucFirstName + "(" + dataType + " $" + key + "): self\n")
		builder.WriteString("    {\n")
		builder.WriteString("        return $this->" + key + " = $" + key + ";\n")
		builder.WriteString("        return $this;\n")
		builder.WriteString("    }\n")
		builder.WriteString("\n")

		builder.WriteString("    public function get" + ucFirstName + "(): " + dataType + "\n")
		builder.WriteString("    {\n")
		builder.WriteString("        return $this->" + key + ";\n")
		builder.WriteString("    }\n")
		builder.WriteString("\n")
	}
	builder.WriteString("}")

	return builder.String()
}
