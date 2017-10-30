package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "gonne",
	Run: func(cmd *cobra.Command, args []string) {
		println("gonne is my ai friend")
	},
}
