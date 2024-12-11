package server

import (
	"database/sql"
	"fmt"
	"musicRecsServer/database"
	"net/http"
)

func endpointAddToBlacklist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		artist := r.URL.Query().Get("artist")

		machineID, err := getMachineID()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}

		err = database.AddToBlacklist(db, machineID, artist)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
		}
	}

}
