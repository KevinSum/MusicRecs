package cli

import "fmt"

func commandHelp(args ...interface{}) error {
	fmt.Println()
	for key, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", key, cmd.description)
	}
	fmt.Println()
	return nil
}
