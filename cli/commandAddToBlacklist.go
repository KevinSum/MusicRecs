package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func commandAddToBlacklist(args ...interface{}) error {
	if len(args) != 1 {
		return fmt.Errorf("getSimilarArtists expects one argument: the artist name")
	}

	// Extract arguments
	artist, ok := args[0].(string)

	// Validate argument types
	if !ok {
		return fmt.Errorf("invalid argument types: expected (string)")
	}

	// Construct url
	url := fmt.Sprintf("http://localhost:8080/addToBlacklist?artist=%s", url.QueryEscape(artist))

	// Create a new HTTP GET request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	fmt.Println(artist + " added to blacklist")

	return nil
}
