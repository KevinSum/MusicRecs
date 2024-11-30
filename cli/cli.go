package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	description string
	callback    func() error
}

func StartCLI() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Get input from user
		fmt.Print("MusicRec > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words

		// Check if we have a command for the input. If so call the callback function
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println("Error calling command:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"hello": {
			description: "Say hello",
			callback:    commandHello,
		},
	}
}

func cleanInput(text string) string {
	output := strings.ToLower(text)
	return output
}
