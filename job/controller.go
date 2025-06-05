package job

import (
	"net/http"
)

func (s *Scope) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the application"))
}

func (s *Scope) Register(w http.ResponseWriter, r *http.Request) {
}

func (s *Scope) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}
