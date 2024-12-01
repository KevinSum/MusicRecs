package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const similarArtistBaseURL = "https://ws.audioscrobbler.com/2.0/?method=artist.getsimilar"

// Struct to parse JSON response into when getting similar artists from last.FM API

type similarArtistsData struct {
	SimilarArtists struct {
		Artist []struct {
			Name string `json:"name"`
		} `json:"artist"`
	} `json:"similarartists"`
}

func getSimilarArtists(artist string, similarArtistsData *similarArtistsData) error {
	// Construct our endpoint URL
	url := similarArtistBaseURL + fmt.Sprintf("&artist=%s&api_key=%s&autocorrect=1&format=json",
		artist, apiKey)

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
	err = json.Unmarshal(body, &similarArtistsData)
	if err != nil {
		return err
	}

	return nil
}
