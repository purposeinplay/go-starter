package main

import (
	"fmt"

	"github.com/oakeshq/go-starter/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
