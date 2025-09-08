package main

import (
	"fmt"
	"io"	// provides interfaces and functions for reading, writing to various sources
	"log"
	"net/http"
	"strings" // for splitting URL path
)

func main() {	
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world\n")
		fmt.Println(w.Header()) //able to print only header
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
	// start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/* Notes on Go http components*/

//  http.ListenAndServe() --> listens on specified TCP network address
			// --> uses a handler to process incoming HTTP requests
			// --> starts an HTTP server and runs indefinitely until shutdown
			// --> forces developers to explicitly handle the server's termination
			
// http.HandleFunc() --> allows you to associate specific functions with different web server endpoints

// http.ResponseWriter --> is an interface that specifies 3 methods
			// --> Header(), returns header map that will be sent
			// --> write([]byte) () - writes data to the connection as part of HTTP reply
			// --> WriteHeader(statusCode int) - Sends HTTP response header w/ status code
			// --> Therefore 'w' is the response construction mechanism
// *http.Request --> pointer to an http.Request struct, containing method, URL path, headers, body

// io.WriteString() --> send text data as part of an HTTP response back to the client

// log.Println() vs fmt.Println() --> both will show messages on Terminal window
				// --> but log.Println() adds timestamp and other prefixes
				
// log.Fatal() --> log error message and terminate with Exit(1)
