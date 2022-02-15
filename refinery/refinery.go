package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Set some port variables
	const port = "9800"
	const address = ":" + port
	// Setup functions to handle various URI's
	http.HandleFunc("/", rootHandler)
	// Start a server
	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	io.WriteString(w, "So refined\n")
}
