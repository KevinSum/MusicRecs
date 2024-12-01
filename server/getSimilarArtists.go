package server

import (
	"MusicRecs/server/lastFM_API"
	"encoding/json"
	"fmt"
	"net/http"
)

func getSimilarArtists(w http.ResponseWriter, r *http.Request) {
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
