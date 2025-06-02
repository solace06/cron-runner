package api

import "net/http"

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Details  string `json:"details,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func BadRequest(detail, instance string) *ProblemDetails {
	return &ProblemDetails{
		Type:     "/problem/bad-request",
		Title:    "Bad Request",
		Status:   http.StatusBadRequest,
		Details:  detail,
		Instance: instance,
	}
}

func InvalidJSON(detail, instance string) *ProblemDetails {
	return &ProblemDetails{
		Type:     "/problem/invalid-json",
		Title:    "Invalid JSON",
		Status:   http.StatusBadRequest,
		Details:  detail,
		Instance: instance,
	}
}

func Conflict(detail, instance string) *ProblemDetails {
	return &ProblemDetails{
		Type:     "/problem/conflict",
		Title:    "Conflict",
		Status:   http.StatusConflict,
		Details:  detail,
		Instance: instance,
	}
}

func Internal(detail, instance string) *ProblemDetails {
	return &ProblemDetails{
		Type:     "/problem/internal",
		Title:    "Internal Server Error",
		Status:   http.StatusInternalServerError,
		Details:  detail,
		Instance: instance,
	}
}