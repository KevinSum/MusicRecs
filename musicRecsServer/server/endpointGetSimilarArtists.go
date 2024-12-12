package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"musicRecsServer/database"
	"musicRecsServer/lastFM_API"
	"net/http"
)

func endpointGetSimilarArtists(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		artist := r.URL.Query().Get("artist")

		machineID, err := getMachineID()
		if err != nil {
			fmt.Printf("Error getting machine ID: %+v\n", err)
		}

		// Call lastFM API to get similar artists
		var similarArtistsData lastFM_API.SimilarArtistsData
		err = lastFM_API.GetSimilarArtists(artist, &similarArtistsData)
		if err != nil {
			fmt.Printf("Error getting similar artists from lastFM API: %+v\n", err)
		}

		// Create list of 10 most similar artists that aren't blacklisted
		var artists []string
		for _, artist := range similarArtistsData.SimilarArtists.Artist {
			var isBlacklisted bool
			isBlacklisted, err = database.IsBlacklisted(db, machineID, artist.Name)
			if err != nil {
				fmt.Printf("Error checking if artist is blacklisted: %+v\n", err)
				return
			}
			if !isBlacklisted {
				artists = append(artists, artist.Name)
				if len(artists) >= numResults {
					break
				}
			}
		}

		// Encode into JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists)
	}
}
