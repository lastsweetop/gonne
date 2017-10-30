package cmd

import (
	"io/ioutil"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Gonne",
	Run: func(cmd *cobra.Command, args []string) {
		strb, _ := ioutil.ReadFile("gonne.lock")
		command := exec.Command("kill", string(strb))
		command.Start()
		println("gonne stop")
	},
}
