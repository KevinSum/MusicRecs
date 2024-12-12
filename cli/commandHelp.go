package main

import "fmt"

func commandHelp(args ...interface{}) error {
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("Command: %s\n", cmd.usage)
		fmt.Printf("Description: %s\n\n", cmd.description)
	}
	fmt.Println()
	return nil
}
