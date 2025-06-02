package api

import (
	"encoding/json"
	"net/http"
)

func WriteProblem(w http.ResponseWriter, pd *ProblemDetails){
	w.Header().Set("Content-Type","application/problem+json")
	w.WriteHeader(pd.Status)
	json.NewEncoder(w).Encode(pd)
}

func WriteResponse(w http.ResponseWriter, status int, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}