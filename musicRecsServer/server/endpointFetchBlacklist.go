package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"musicRecsServer/database"
	"net/http"
)

func endpointFetchBlacklist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		machineID, err := getMachineID()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

		artists, err := database.FetchBlacklist(db, machineID)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists)
	}

}
