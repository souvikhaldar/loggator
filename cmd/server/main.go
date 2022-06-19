package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func NewServer() *server {
	return &server{
		router: mux.NewRouter(),
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// handleLogsPost handles logs as a string and stores it
func (s *server) handleLogsPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
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

func main() {
	s := NewServer()
	s.router.HandleFunc("/logs/", s.handleLogsPost()).Methods("POST")
	s.router.HandleFunc("/logs/", s.handleLogsGet()).Methods("GET")
	log.Fatal(http.ListenAndServe(":8192", s))
}
