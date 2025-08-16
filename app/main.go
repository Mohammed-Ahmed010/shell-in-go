package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var allCommands = []string{"exit", "type", "echo", "pwd"}

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
		case "pwd":
			handlePWD()
		case "cd":
			handleCD(args)
		default:
			{
				_, exits := findBinInPath(command)
				if exits {
					cmd := exec.Command(command, args...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Run()
				} else {
					fmt.Println(command + ": command not found")

				}
			}

		}

	}

}

func main() {
	run()

}
