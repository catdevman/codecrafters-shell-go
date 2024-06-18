package main

import (
	"bufio"
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

func main() {
	allowedCommands := []string{}
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	found := false
	for _, cmd := range allowedCommands {
		if cmd == command {
			found = true
		}
	}
	if !found {
		fmt.Sprintf("%s: command not found", command)
	}
}
