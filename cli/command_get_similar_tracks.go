package cli

import (
	"MusicRecs/server/lastFM_API"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandGetSimilarTracks(args ...interface{}) error {
	//fmt.Printf("Args received: %v\n", args)
	//fmt.Printf("Arg[0]: %v (type: %T)\n", args[0], args[0])
	//fmt.Printf("Arg[1]: %v (type: %T)\n", args[1], args[1])

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
	var similarTracksData lastFM_API.SimilarTracksData
	err = json.Unmarshal(body, &similarTracksData)
	if err != nil {
		return err
	}

	// Print the results
	for _, track := range similarTracksData.SimilarTracks.Track {
		fmt.Println(track.Name + " (" + track.Artist.Name + ")")
	}

	return nil
}
