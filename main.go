package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "os/exec"
    "strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)

	var history []string // slice to store input history

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		// Check for errors
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// save user history
		history = append(history, input)

		// execute input
		if err = execInput(input, history); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}


func execInput(input string, history []string) error {
	// prepare command and args for os
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	// check for commands not handled by os
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "history":
		for i, cmd := range history {
			fmt.Printf("%d: %s", i+1, cmd)
		}
		return nil
	case "exit":
		os.Exit(0)
	}

	// configure command
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// run and return command
	return cmd.Run()
}