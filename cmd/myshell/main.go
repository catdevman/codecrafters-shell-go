package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	allowedCommands := []string{"exit", "echo", "type", "pwd"}
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
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", c)
				} else {
					f := false
					paths := strings.Split(os.Getenv("PATH"), ":")
					for _, path := range paths {
						fp := filepath.Join(path, c)

						if _, err := os.Stat(fp); err == nil {

							fmt.Println(fp)
							f = true

						}
					}
					if !f {
						fmt.Fprintf(os.Stdout, "%s: not found\n", c)
					}
				}
				break
			case "pwd":
				cwd, err := os.Getwd()
				if err != nil {
					//TODO: Should probably do something here
				}
				fmt.Println(cwd)
				break
			}
		} else {
			f := false
			paths := strings.Split(os.Getenv("PATH"), ":")
			for _, path := range paths {
				fp := filepath.Join(path, cmd)

				if _, err := os.Stat(fp); err == nil {
					f = true
					command := exec.Command(cmd, cmdPieces[1:]...)
					command.Stdout = os.Stdout
					command.Stderr = os.Stderr

					err := command.Run()

					if err != nil {
						fmt.Fprintf(os.Stdout, "%s: not found\n", cmd)
					}
				}
			}

			if !f {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
			}
		}
	}
}
