package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vumanskyi/data-object-convert/pkg/builder"
	"github.com/vumanskyi/data-object-convert/pkg/convert"
	"github.com/vumanskyi/data-object-convert/pkg/export"
)

var (
	jsonCmd = &cobra.Command{
		Use:   "json",
		Short: "Convert JSON object into Data Object",
		Long:  "Convert JSON object into Data Object",
		Run: func(cmd *cobra.Command, args []string) {
			mockStr := "{\"name\": \"Vlad\", \"age\": 30}"
			fmt.Println(mockStr)
			data, err := parser.Parse(cmd.Use, mockStr)
			if nil != err {
				fmt.Println("Error on parsing json", err)
				return
			}

			factory := builder.NewObjectBuilderFactory()
			b, err := factory.Create("dto")

			if nil != err {
				fmt.Println("Error on create builder", err)
				return
			}

			e := export.NewDataExport("stdout")
			convert.NewObjectGenerator(e, b).Generate(data)
		},
	}
)

func init() {
	rootCmd.AddCommand(jsonCmd)

	//jsonCmd.Flags().String("")
}
