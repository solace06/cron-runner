package api

import "net/http"

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Details  string `json:"details,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func BadRequest(detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "/problem/bad-request",
		Title:    "Bad Request",
		Status:   http.StatusBadRequest,
		Details:  detail,
		Instance: instance,
	}
}

func InvalidJSON(detail, instance string) ProblemDetails {
	return ProblemDetails{
		Type:     "/problem/invalid-json",
		Title:    "Invalid JSON",
		Status:   http.StatusBadRequest,
		Details:  detail,
		Instance: instance,
	}
}
