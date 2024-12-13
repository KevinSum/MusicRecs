package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandGetBlacklist(args ...interface{}) error {
	// Construct url
	url := baseURL + port + "/fetchBlacklist"

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
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

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse the JSON response
	var blacklist []string
	err = json.Unmarshal(body, &blacklist)
	if err != nil {
		return err
	}

	// Print the results
	for _, artist := range blacklist {
		fmt.Println(artist)
	}
	fmt.Println()

	return nil
}
