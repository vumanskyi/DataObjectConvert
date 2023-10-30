package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vumanskyi/data-object-convert/pkg/context"
)

var (
	jsonCmd = &cobra.Command{
		Use:   "json",
		Short: "Convert JSON object into Data Object",
		Long:  "Convert JSON object into Data Object",
		Run:   context.Run,
	}
)

func init() {
	rootCmd.AddCommand(jsonCmd)
}
