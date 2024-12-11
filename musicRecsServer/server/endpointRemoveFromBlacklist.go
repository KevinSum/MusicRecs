package server

import (
	"database/sql"
	"fmt"
	"musicRecsServer/database"
	"net/http"
)

func endpointRemoveFromBlacklist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		artist := r.URL.Query().Get("artist")

		machineID, err := getMachineID()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

		// TODO: Check if it's actually in blacklist

		err = database.RemoveFromBlacklist(db, machineID, artist)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}
	}

}
