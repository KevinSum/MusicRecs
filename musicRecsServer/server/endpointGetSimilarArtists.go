package server

import (
	"encoding/json"
	"fmt"
	"musicRecsServer/lastFM_API"
	"net/http"
)

func endpointGetSimilarArtists(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	artist := r.URL.Query().Get("artist")

	// Call lastFM API to get similar artists
	var similarArtistsData lastFM_API.SimilarArtistsData
	err := lastFM_API.GetSimilarArtists(artist, &similarArtistsData)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	// TODO: Do some filtering here later

	// Encode into JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(similarArtistsData)
}
