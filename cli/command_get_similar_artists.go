package cli

import (
	"MusicRecs/server/lastFM_API"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandGetSimilarArtists(args ...interface{}) error {
	// Extract arguments
	artist, ok := args[0].(string)

	// Validate argument types
	if !ok {
		return fmt.Errorf("invalid argument types: expected (string)")
	}

	// Construct url
	url := fmt.Sprintf("http://localhost:8080/getSimilarArtists?artist=%s", artist)

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
	var similarTracksData lastFM_API.SimilarArtistsData
	err = json.Unmarshal(body, &similarTracksData)
	if err != nil {
		return err
	}

	// Print the results
	for _, track := range similarTracksData.SimilarArtists.Artist {
		fmt.Println(track.Name)
	}

	return nil
}
