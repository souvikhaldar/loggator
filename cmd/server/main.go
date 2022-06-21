package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souvikhaldar/loggator/pkg/db"
	"github.com/souvikhaldar/loggator/pkg/logs"
)

type server struct {
	router *mux.Router
	logger logs.Repository
}

func NewServer() *server {
	return &server{
		router: mux.NewRouter(),
		logger: db.NewDB(),
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// handleLogsPost handles logs as a string and stores it
func (s *server) handleLogsPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ld := new(logs.LogData)
		if err := json.NewDecoder(r.Body).Decode(ld); err != nil {
			err = fmt.Errorf("Error in parsing payload: %s", err)
			log.Println(err)
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
		if err := s.logger.StoreLog(*ld); err != nil {
			err = fmt.Errorf("Error in storing data: %s", err)
			log.Println(err)
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

// handleLogsGet returns logs based on filters
func (s *server) handleLogsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filters := r.URL.Query()
		log.Printf("Filters: %+v\n", filters)
		logs, err := s.logger.FetchLog(filters)
		if err != nil {
			err = fmt.Errorf("Error in fetching logs: %s", err)
			log.Println(err)
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
		}
		log.Println("Logs:")
		for _, l := range logs {
			log.Printf("%+v\n", l)
		}
		if err := json.NewEncoder(w).Encode(logs); err != nil {
			err = fmt.Errorf("Error in encoding json: %s", err)
			log.Println(err)
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

func main() {
	s := NewServer()
	s.router.HandleFunc("/logs", s.handleLogsPost()).Methods("POST")
	s.router.HandleFunc("/logs", s.handleLogsGet()).Methods("GET")
	log.Fatal(http.ListenAndServe(":8192", s))
}
