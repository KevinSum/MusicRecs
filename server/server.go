package server

import (
	"fmt"
	"log"
	"net/http"
)

func SetupServer(done chan bool) {
	const filepathRoot = "."
	const port = "8080"

	// Setup HTTP request multiplexer
	mux := http.NewServeMux()

	// Register handler (we a create a file server that servers from a local file system)
	// for a given URL pattern (/ = URL rootpath)
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// Add endpoints for functions
	mux.HandleFunc("/test", test)

	// Set up HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)

	// Signal to channel that server is set up
	done <- true

	// Start listening and serving
	log.Fatal(srv.ListenAndServe())
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World!")
}
