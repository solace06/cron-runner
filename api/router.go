package api

import (
	"github.com/gorilla/mux"
	"github.com/solace06/cron-runner/job"
)

func NewRouter(s *job.Scope) *mux.Router {
	r := mux.NewRouter()

	//public route
	r.HandleFunc("/", s.Home).Methods("GET")
	r.HandleFunc("/login", s.Login).Methods("POST")

	//protected routes

	return r
}
