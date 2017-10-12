package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gonne",
	Run: func(cmd *cobra.Command, args []string) {
		println("gonne version is 0.0.1")
	},
}
