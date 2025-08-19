package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func handleType(args []string) {
	if len(args) == 0 {
		fmt.Println("missing command! enter command  <[command]>")
		return
	}
	if slices.Contains(allCommands, args[0]) {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return

	}
	file, exists := findBinInPath(args[0])
	if exists {
		fmt.Fprintf(os.Stdout, "%s is %s\n", args[0], file)
		return
	}
	fmt.Printf("%s: not found \n", args[0])

}

func handlePWD() {
	dr, err := os.Getwd()
	if err != nil {
		fmt.Println("error printing working directory", err)
	}
	fmt.Println(dr)
}

func handleEcho(args []string) {
	text := strings.Join(args, " ")
	fmt.Println(text)
}
func handleCD(args []string) {
	route := args[0]
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error changing directory")
	}
	if route[0] == '~' {
		route = home + route[1:]
	}
	err = os.Chdir(route)
	if err != nil {
		fmt.Println("error changing directory")
	} else {
		fmt.Println(route)
	}

}
func findBinInPath(bin string) (string, bool) {
	paths := os.Getenv("PATH")
	for _, path := range strings.Split(paths, ":") {
		file := path + "/" + bin
		info, err := os.Stat(file)
		if err == nil && (info.Mode()&0111) != 0 {
			return file, true
		}
	}

	return "", false
}
