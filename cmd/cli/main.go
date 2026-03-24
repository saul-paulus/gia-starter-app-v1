package main

import (
	"fmt"
	"gia-starter-app-V1/internal/cli"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: make:<command> <name>")
		return
	}
	cli.Run(os.Args)
}
