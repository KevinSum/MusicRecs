package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type callback func(args ...interface{}) error

type cliCommand struct {
	description string
	callback    callback
}

func startCLI() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Get input from user
		fmt.Print("MusicRec > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}

		// Split the input into command and arguments
		parts := strings.Fields(input)
		commandName := parts[0] // The command is the first part

		// Get the rest as arguments
		var args []interface{}
		if len(parts) > 1 {
			for _, arg := range parts[1:] {
				args = append(args, arg)
			}
		}

		// Check if we have a command for the input. If so call the callback function
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(args...)
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
		"help": {
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"getSimilarArtists": {
			description: "Retrive list of similar artists",
			callback:    commandGetSimilarArtists,
		},
		"getSimilarTracks": {
			description: "Retrive list of similar tracks",
			callback:    commandGetSimilarTracks,
		},
	}
}
