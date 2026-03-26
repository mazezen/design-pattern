package main

import (
	"fmt"
	"os"

	command "github.com/mazezen/design-pattern/go/15_command"
)

/**
* go run main.go hello
* # hello world

* go run main.go hello mazezen
* # hello mazezen

* go run main.go add 3 5
* # 8
*
* go run main.go sub
* #
* unknown command
 */
func main() {

	commands := []command.Command{
		&command.HelloCommand{},
		&command.AddCommand{},
	}

	if len(os.Args) < 2 {
		fmt.Println("usage: cli <command>")
		return
	}

	cmdName := os.Args[1]
	args := os.Args[2:]
	fmt.Printf("commnd name: %s\n", cmdName)
	fmt.Printf("args: %v\n", args)

	for _, cmd := range commands {
		if cmd.Name() == cmdName {
			result := cmd.Execute(args)
			fmt.Println(result)
			return
		}
	}

	fmt.Println("unknown command")
}
