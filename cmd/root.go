package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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

func init() {
	rootCmd.PersistentFlags().StringP("class", "c", "", "Class name.")
	rootCmd.PersistentFlags().StringP("type", "t", "dto", "Convert to object type.")
	rootCmd.PersistentFlags().StringP("export", "e", "stdout", "Export result.")
	rootCmd.MarkPersistentFlagRequired("class")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to run command: %v\n\n", err)
		os.Exit(1)
	}
}
