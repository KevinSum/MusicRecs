package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type similarArtistsData struct {
	SimilarArtists struct {
		Artist []struct {
			Name string `json:"name"`
		} `json:"artist"`
	} `json:"similarartists"`
}

func commandGetSimilarArtists(args ...interface{}) error {
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

	url := fmt.Sprintf(baseURL+port+"/getSimilarArtists?artist=%s", artist)

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
	var similarArtists []string
	err = json.Unmarshal(body, &similarArtists)
	if err != nil {
		return err
	}

	// Print the results
	for _, artist := range similarArtists {
		fmt.Println(artist)
	}
	fmt.Println()

	return nil
}
