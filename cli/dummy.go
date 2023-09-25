package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	woke "github.com/mahmoudgalalz/woke"
)

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func main() {
	cliFunc := func() error {
		// name := StringPrompt("What is your name?")
		// fmt.Println("Hello ", name)

		// gh CLI test
		cmd := exec.Command("gh", "repo", "create")

		// Set the command to use the current process's stdin, stdout, and stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Execute the command
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running 'gh' command:", err)
			os.Exit(1)
		}
		return nil
	}

	woke.Wrap(cliFunc)
}
