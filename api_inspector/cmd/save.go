package cmd

import (
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [name] [url]",
	Short: "Save a request configuration",
	Args:  cobra.ExactArgs(2),
	Run:   saveRequestConfig,
}
