package main

import (
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

func main() {
	srv := NewServer()
	srv.routes()
	log.Fatal(http.ListenAndServe(":8192", srv))
}
