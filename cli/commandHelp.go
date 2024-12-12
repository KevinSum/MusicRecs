package main

import "fmt"

func commandHelp(args ...interface{}) error {
	for _, cmd := range getCommands() {
		fmt.Printf("Command: %s\n", cmd.usage)
		fmt.Printf("Description: %s\n\n", cmd.description)
	}
	return nil
}
