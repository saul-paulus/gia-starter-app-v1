package cli

import "fmt"

func Run(args []string) {
	command := args[1]

	switch command {
	case "make:module":
		MakeModule(args)
	default:
		fmt.Println("Unknown command:", command)
	}
}
