package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vumanskyi/data-object-convert/pkg/context"
)

var (
	xmlCmd = &cobra.Command{
		Use:   "xml",
		Short: "Convert XML object into Data Object",
		Long:  "Convert XML object into Data Object",
		Run:   context.Run,
	}
)

func init() {
	rootCmd.AddCommand(xmlCmd)
}
