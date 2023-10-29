package convert

import (
	"github.com/vumanskyi/data-object-convert/pkg/builder"
	"github.com/vumanskyi/data-object-convert/pkg/export"
)

type Generator interface {
	Generate(data interface{})
}

// ObjectGenerator - generator struct
type ObjectGenerator struct {
	exporter export.Exporter
	builder  builder.Builder
}

func NewObjectGenerator(e export.Exporter, b builder.Builder) *ObjectGenerator {
	return &ObjectGenerator{
		e,
		b,
	}
}

func (og *ObjectGenerator) Generate(data map[string]interface{}) {
	og.exporter.Export(og.builder.Build(data))
}
