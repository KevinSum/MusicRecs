package lastFM_API

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const similarTrackBaseURL = "https://ws.audioscrobbler.com/2.0/?method=track.getsimilar"

// Struct to parse JSON response into when getting similar tracks from last.FM API
type SimilarTracksData struct {
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

func GetSimilarTracks(track, artist string, similarTracksData *SimilarTracksData) error {
	// Construct our endpoint URL
	url := similarTrackBaseURL + fmt.Sprintf("&artist=%s&track=%s&api_key=%s&autocorrect=1&format=json",
		artist, track, apiKey)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response JSON
	err = json.Unmarshal(body, &similarTracksData)
	if err != nil {
		return err
	}

	return nil
}
