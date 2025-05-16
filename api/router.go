package api

import (
	"github.com/gorilla/mux"
	"github.com/solace06/cron-runner/job"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	//public route
	r.HandleFunc("/", job.Home).Methods("GET")
	r.HandleFunc("/login", job.Login).Methods("POST")

	//protected routes
	return r
}
