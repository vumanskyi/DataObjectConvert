package context

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vumanskyi/data-object-convert/pkg/builder"
	"github.com/vumanskyi/data-object-convert/pkg/convert"
	"github.com/vumanskyi/data-object-convert/pkg/export"
	baseParser "github.com/vumanskyi/data-object-convert/pkg/parser"
	"log"
)

func Run(cmd *cobra.Command, args []string) {
	if 0 == len(args) {
		log.Fatal("Missed input data")
	}
	inputData := args[0]

	className, _ := cmd.Flags().GetString("class")
	objectType, _ := cmd.Flags().GetString("type")
	exportType, _ := cmd.Flags().GetString("export")

	var parser = baseParser.NewStructParser()
	data, err := parser.Parse(cmd.Use, inputData)
	if nil != err {
		fmt.Println("Error on parsing json", err)
		return
	}

	factory := builder.NewObjectBuilderFactory()
	b, err := factory.Create(objectType)

	if nil != err {
		fmt.Println("Error on create builder", err)
		return
	}

	e := export.NewDataExport(exportType)
	convert.NewObjectGenerator(e, b).Generate(className, data)
}
