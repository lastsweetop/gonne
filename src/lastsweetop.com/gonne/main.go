package main

import (
	"fmt"
	"os"

	"lastsweetop.com/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
