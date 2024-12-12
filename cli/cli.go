package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type callback func(args ...interface{}) error

type cliCommand struct {
	usage       string
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
		parts := strings.SplitN(input, " ", 2)
		commandName := parts[0] // The command is the first part

		// Split the arguments by comma
		var args []interface{}
		if len(parts) > 1 {
			argParts := strings.SplitN(parts[1], ", ", -1)
			for _, arg := range argParts[0:] {
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
			usage:       "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"getSimilarArtists": {
			usage:       "getSimilarArtists 'artist'",
			description: "Retrive list of similar artists",
			callback:    commandGetSimilarArtists,
		},
		"getSimilarTracks": {
			usage:       "getSimilarTracks 'track', 'artist'",
			description: "Retrive list of similar tracks",
			callback:    commandGetSimilarTracks,
		},
		"addToBlacklist": {
			usage:       "addToBlacklist 'artist'",
			description: "Add an artist to blacklist so that they won't be recommended",
			callback:    commandAddToBlacklist,
		},
		"removeFromBlacklist": {
			usage:       "removeFromBlacklist 'artist'",
			description: "Remove artist from blacklist",
			callback:    commandRemoveFromBlacklist,
		},
		"getBlacklist": {
			usage:       "getBlacklist",
			description: "Retrive list of blacklisted artists",
			callback:    commandGetBlacklist,
		},
	}
}
