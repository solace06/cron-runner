package api

import (
	"github.com/gorilla/mux"
	"github.com/solace06/cron-runner/job"
)

func NewRouter(s *job.Scope) *mux.Router {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()

	//public route
	r.HandleFunc("/", s.Home).Methods("GET")
	v1.HandleFunc("/register", s.Register).Methods("POST")
	r.HandleFunc("/login", s.Login).Methods("POST")

	//protected routes

	return r
}
