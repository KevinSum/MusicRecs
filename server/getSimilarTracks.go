package server

import (
	"MusicRecs/server/lastFM_API"
	"encoding/json"
	"fmt"
	"net/http"
)

func getSimilarTracks(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	artist := r.URL.Query().Get("artist")
	track := r.URL.Query().Get("track")

	// Call lastFM API to get similar tracks
	var similarTracksData lastFM_API.SimilarTracksData
	err := lastFM_API.GetSimilarTracks(track, artist, &similarTracksData)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	// TODO: Do some filtering here later

	// Encode into JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(similarTracksData)
}
