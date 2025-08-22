package cmd

import (
	"github.com/spf13/cobra"
)


var postCmd = &cobra.Command{
	Use:   "post [url]",
	Short: "Make a POST request",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {
		method = "POST"
		makeRequest(cmd, args)
	},
}