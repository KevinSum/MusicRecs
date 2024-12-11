package server

import (
	"fmt"
	"log"
	"musicRecsServer/database"
	"net/http"
	"os/exec"
	"strings"
)

func SetupServer() {
	database, err := database.CreateDatabase()
	if err != nil {
		log.Fatalf("Error creating blacklist table: %v", err)
	}

	// Setup HTTP request multiplexer
	mux := http.NewServeMux()

	// Register handler (we a create a file server that servers from a local file system)
	// for a given URL pattern (/ = URL rootpath)
	const filepathRoot = "."
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// Add endpoints for functions
	mux.HandleFunc("/getSimilarTracks", endpointGetSimilarTracks)
	mux.HandleFunc("/getSimilarArtists", endpointGetSimilarArtists)
	mux.HandleFunc("/addToBlacklist", endpointAddToBlacklist(database))
	mux.HandleFunc("/removeFromBlacklist", endpointRemoveFromBlacklist(database))
	mux.HandleFunc("/fetchBlacklist", endpointFetchBlacklist(database))

	// Set up HTTP server
	const port = "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)

	// Start listening and serving
	log.Fatal(srv.ListenAndServe())
}

func getMachineID() (string, error) {
	// Get machine ID to use to get artist blacklist for this machine
	cmd := exec.Command("sh", "-c", "cat /sys/class/net/*/address | head -n 1")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get machine ID: %v", err)
	}
	machineID := strings.TrimSpace(string(output))
	return machineID, nil
}
