package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var allCommands = []string{"exit", "type", "echo"}

func hanldeInput() (command string, args []string) {
	cliInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	cliTrimmed := strings.TrimSpace(cliInput)
	cliArray := strings.Split(cliTrimmed, " ")
	command, args = cliArray[0], cliArray[1:]
	return command, args

}
func run() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, args := hanldeInput()

		switch command {
		case "exit":
			// can check  here for exit code
			os.Exit(0)

		case "echo":
			handleEcho(args)
		case "type":

			handleType(args)

		default:
			{
				fmt.Println(command + ": command not found")
			}

		}

	}

}

func main() {
	run()

}
