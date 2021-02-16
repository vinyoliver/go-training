package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type personRequest struct {
	Name string
	Age  int
}

type personResponse struct {
	Message string
}

func main() {

	// Create a mux for routing incoming requests
	m := http.NewServeMux()

	// All URLs will be handled by this function
	m.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			response := &personResponse{
				Message: "method not allowed",
			}
			// and encode our response to JSON while writing to the ResponseWriter
			json.NewEncoder(w).Encode(&response)
			return
		}

		var payload personRequest
		// here we're decoding the request body from JSON to our request struct
		json.NewDecoder(r.Body).Decode(&payload)

		// setting the content-type response header
		w.Header().Set("Content-Type", "application/json")
		// setting our response status code
		w.WriteHeader(http.StatusOK)
		response := &personResponse{
			Message: fmt.Sprintf("My name is %v and I'm %v years old!", payload.Name, payload.Age),
		}
		// encoding our response to JSON and writing to the responseWriter
		json.NewEncoder(w).Encode(response)
	})

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// Create a server listening on port 8000
	s := &http.Server{
		Addr:    ":8000",
		Handler: m,
	}

	// Continue to process new requests until an error occurs
	log.Fatal(s.ListenAndServe())
}
