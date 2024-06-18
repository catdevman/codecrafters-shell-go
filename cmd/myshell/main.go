package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	allowedCommands := []string{"exit", "echo", "type"}
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fullCmd := strings.TrimRight(input, "\n")
		cmdPieces := strings.Split(fullCmd, " ")
		cmd := cmdPieces[0]

		if slices.Contains(allowedCommands, cmd) {
			switch cmd {
			case "exit":
				i, _ := strconv.ParseInt(cmdPieces[1], 32, 10)
				os.Exit(int(i))
				break
			case "echo":
				theRest := cmdPieces[1:]
				fmt.Fprint(os.Stdout, strings.Join(theRest, " ")+"\n")
				break
			case "type":
				c := cmdPieces[1]
				if slices.Contains(allowedCommands, c) {
					fmt.Fprintf(os.Stdout, "%s is a shell builtin", c)
				} else {
					fmt.Fprintf(os.Stdout, "%s: not found", c)
				}
				break
			}
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}
