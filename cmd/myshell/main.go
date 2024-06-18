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
	allowedCommands := []string{"exit"}
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fullCmd := strings.TrimRight(input, "\n")
		cmdPieces := strings.Split(fullCmd, " ")
		cmd := cmdPieces[0]

		if slices.Contains(allowedCommands, cmd) {

			if cmd == "exit" {
				i, _ := strconv.ParseInt(cmdPieces[1], 32, 10)
				os.Exit(int(i))
			}
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}
