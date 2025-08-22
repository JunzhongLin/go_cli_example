package cmd

import (
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load [name]", 
	Short: "Load and execute a saved request",
	Args:  cobra.ExactArgs(1),
	Run:   loadRequestConfig,
}