package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vumanskyi/data-object-convert/pkg/context"
)

var (
	yamlCmd = &cobra.Command{
		Use:   "yaml",
		Short: "Convert YAML object into Data Object",
		Long:  "Convert YAML object into Data Object",
		Run:   context.Run,
	}
)

func init() {
	rootCmd.AddCommand(yamlCmd)
}
