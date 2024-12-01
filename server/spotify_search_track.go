package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Struct to parse JSON response into when searching for a track using Spotify API
type trackInfo struct {
	Tracks struct {
		Items []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
		} `json:"items"`
	} `json:"tracks"`
}

func getTrackID(searchInput string, accessToken string, trackInfo *trackInfo) error {
	// Embed search input into our URL.
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=1",
		url.QueryEscape(searchInput))

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response JSON
	err = json.Unmarshal(body, &trackInfo)
	if err != nil {
		return err
	}

	return nil
}
