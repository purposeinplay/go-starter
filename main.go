package main

import (
	"fmt"

	"github.com/purposeinplay/go-starter/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
