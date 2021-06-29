package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize playlist directory name and port #
	const playlistDir = "playlists"
	const port = 8080

	// Handler for the files
	http.Handle("/", enableHeaders(http.FileServer(http.Dir(playlistDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", playlistDir, port)

	// Log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// enableHeaders acts as middleware to enable CORS headers
func enableHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}