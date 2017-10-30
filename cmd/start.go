package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	var daemon bool

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start Gonne",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("gonne", "start")
				command.Start()
				fmt.Printf("gonne start, [PID] %d running...\n", command.Process.Pid)
				ioutil.WriteFile("gonne.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
				daemon = false
				os.Exit(0)
			} else {
				fmt.Println("gonne start")
			}
			startHttp()
		},
	}
	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	RootCmd.AddCommand(startCmd)

}

func startHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello cmd!")
	})
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
