package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct to parse JSON response into when getting similar tracks from last.FM API
type similarTracksData struct {
	SimilarTracks struct {
		Track []struct {
			Name   string `json:"name"`
			Artist struct {
				Name string `json:"name"`
			} `json:"artist"`
			URL string `json:"url"`
		} `json:"track"`
	} `json:"similartracks"`
}

func commandGetSimilarTracks(args ...interface{}) error {
	if len(args) != 2 {
		return fmt.Errorf("getSimilarTracks expects two argument: the track name and artist name")
	}

	// Extract arguments
	track, ok := args[0].(string)
	artist, ok2 := args[1].(string)

	// Validate argument types
	if !ok || !ok2 {
		return fmt.Errorf("invalid argument types: expected (string, string)")
	}

	// Construct url
	url := fmt.Sprintf("http://localhost:8080/getSimilarTracks?track=%s&artist=%s", track, artist)

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
	var similarTracks []string
	err = json.Unmarshal(body, &similarTracks)
	if err != nil {
		return err
	}

	// Print the results
	for _, track := range similarTracks {
		fmt.Println(track)
	}
	fmt.Println()

	return nil
}
