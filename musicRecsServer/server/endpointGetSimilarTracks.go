package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"musicRecsServer/database"
	"musicRecsServer/lastFM_API"
	"net/http"
)

func endpointGetSimilarTracks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		artist := r.URL.Query().Get("artist")
		track := r.URL.Query().Get("track")

		machineID, err := getMachineID()
		if err != nil {
			fmt.Printf("Error getting machine ID: %+v\n", err)
		}

		// Call lastFM API to get similar tracks
		var similarTracksData lastFM_API.SimilarTracksData
		err = lastFM_API.GetSimilarTracks(track, artist, &similarTracksData)
		if err != nil {
			fmt.Printf("Error getting similar tracks from lastFM API: %+v\n", err)
		}

		// Create list of 10 most similar tracks where the artists isn't blacklisted
		var tracks []string
		for _, track := range similarTracksData.SimilarTracks.Track {
			var isBlacklisted bool
			isBlacklisted, err = database.IsBlacklisted(db, machineID, track.Artist.Name)
			if err != nil {
				fmt.Printf("Error checking if artist is blacklisted: %+v\n", err)
				return
			}
			if !isBlacklisted {
				tracks = append(tracks, track.Name+" ("+track.Artist.Name+")")
				if len(tracks) >= numResults {
					break
				}
			}
		}

		// Encode into JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tracks)
	}
}
