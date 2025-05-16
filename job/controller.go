package job

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the application"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the application"))
}