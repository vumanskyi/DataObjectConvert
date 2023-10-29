package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	baseParser "github.com/vumanskyi/data-object-convert/pkg/parser"
	"github.com/vumanskyi/data-object-convert/version"
	"os"
)

var (
	className string

	rootCmd = cobra.Command{
		Use:     "data-object-converter",
		Short:   "Convert: Quickly turn JSON | YAML | XML into Data Objects",
		Long:    "Convert: Quickly turn JSON | YAML | XML into Data Objects",
		Version: version.Release,
	}
)

var parser = baseParser.NewStructParser()

//var exporter = export.NewDataExport()
//var convert = baseConvert.NewObjectGenerator(parser)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to run command: %v\n\n", err)
		os.Exit(1)
	}
}
