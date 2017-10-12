package tools

import "fmt"

type Log struct {
}

func (*Log) Info(message string) {
	fmt.Println(message)
}
