package main

import (
	"io"
	"log"
	"net/http"
	"strings" // for splitting URL path
)

func main() {
	// Hello world, the web server
	
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world\n")
	}
	
	goodbyeHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Goodbye, world\n")
	}
	
	// parsing key fields/ methods on http.REquest
	requestHandler := func(w http.ResponseWriter, req *http.Request) {
		// Method (GET, POST, etc.)
		log.Println("Method:", req.Method)
		
		// URL path
		log.Println("Path:", req.URL.Path)
		
		// Query parameters: /hello?name=ALice
		name := "beautiful world" // default
		segments := strings.Split(req.URL.Path, "/")
		if len(segments) > 2 && segments[2] != "" {
			name = segments[2]
		}
		
		// Headers
		log.Println("User-Agent:", req.Header.Get("User-Agent"))
		
		// Request body (for POST/PUT)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}
		log.Println("Body:", string(body))
		
		//respond
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Hello, "+name+"\n")
	}
	
	// HandleFunc()
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)
	http.HandleFunc("/hello/", requestHandler) // use of http.Request
	
	log.Println("Attempting to start server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err == nil {
		log.Printf("Failed to start server: %v\n", err)
	} else {
		log.Fatal(err)
	}
}
