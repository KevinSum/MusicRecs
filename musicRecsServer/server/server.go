package server

import (
	"log"
	"net/http"
)

func SetupServer() {
	// Setup HTTP request multiplexer
	mux := http.NewServeMux()

	// Register handler (we a create a file server that servers from a local file system)
	// for a given URL pattern (/ = URL rootpath)
	const filepathRoot = "."
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// Add endpoints for functions
	mux.HandleFunc("/getSimilarTracks", getSimilarTracksEndpoint)
	mux.HandleFunc("/getSimilarArtists", getSimilarArtistsEndpoint)

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
