package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	// Spotify client ID/Secret obtain from our Spotify App settings
	spotifyClientID     = "b6bbb9c409ff41d4a28774bd45e2ad6d"
	spotifyClientSecret = "3cfbdb4ae7514dafa7375139c31c4890"
)

// Struct to parse JSON response into when requesting access token from Spotify API
type accessToken struct {
	Token               string `json:"access_token"`
	TokenType           string `json:"token_type"`
	ExpiresIn           int    `json:"expires_in"`
	tokenExpirationTime time.Time
}

func getAccessToken(accessToken *accessToken) error {
	// Create data for POST request
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", spotifyClientID)
	data.Set("client_secret", spotifyClientSecret)

	// Create HTTP POST request with data
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token",
		bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	// Set Content-Type header to /x-www-form-urlencoded
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received non-OK HTTP status code: %v", resp.Status)
	}

	// Parse the JSON response into the TokenResponse struct
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return err
	}

	// Set expiration time of this token, so that we know when we need request a new one
	accessToken.tokenExpirationTime = time.Now().Add(
		time.Duration(accessToken.ExpiresIn) * time.Second)

	return nil
}
