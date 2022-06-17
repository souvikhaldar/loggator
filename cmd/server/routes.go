package main

import (
	"log"
	"net/http"
)

// handleLogsPost handles logs as a string and stores it
func (s *server) handleLogsPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		if _, err := r.Body.Read(body); err != nil {
			http.Error(w, "Could not parse the body", http.StatusBadRequest)
			return
		}
		log.Println("Body: ", string(body))
	}
}

// handleLogsGet returns logs based on filters
func (s *server) handleLogsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) routes() {
	s.router.HandleFunc("/logs/", s.handleLogsPost()).Methods("POST")
	s.router.HandleFunc("/logs/", s.handleLogsGet).Methods("GET")
}
