package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vumanskyi/data-object-convert/version"
	"os"
)

var className string

var rootCmd = cobra.Command{
	Use:     "data-object-converter",
	Short:   "Convert: Quickly turn JSON | YAML | XML into Data Objects",
	Long:    "Convert: Quickly turn JSON | YAML | XML into Data Objects",
	Version: version.Release,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to run command: %v\n\n", err)
		os.Exit(1)
	}
}
