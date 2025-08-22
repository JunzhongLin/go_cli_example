package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Make a GET request",
	Args:  cobra.ExactArgs(1), 
	Run:   func(cmd *cobra.Command, args []string) {
		method = "GET"
		makeRequest(cmd, args)
	},
}