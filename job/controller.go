package job

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/solace06/cron-runner/api"
)

func NewRouter(s *Scope) *mux.Router {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()

	//public route
	r.HandleFunc("/", s.Home).Methods("GET")
	v1.HandleFunc("/register", s.Register).Methods("POST")
	r.HandleFunc("/login", s.Login).Methods("POST")

	//protected routes

	return r
}

func (s *Scope) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the application"))
}

func (s *Scope) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		slog.Error("error unmarshalling the request body: ", "error", err.Error())
		api.WriteProblem(w, api.InvalidJSON("Invalid Request Body", ""))
		return
	}

	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	user.UserName = strings.TrimSpace(user.UserName)

	if user.UserName == "" || user.Password == "" || user.Email == "" {
		slog.Error("invalid username, email or password")
		api.WriteProblem(w, api.BadRequest("Username, email or password cannot be empty", ""))
		return
	}
	if len(user.UserName) < 3 {
		slog.Error("invalid username")
		api.WriteProblem(w, api.BadRequest("Username must contain atleast 3 characters", ""))
		return
	}
	if len(user.Password) < 6 {
		slog.Error("invalid password")
		api.WriteProblem(w, api.BadRequest("Password should contain atleast 6 characters", ""))
		return
	}
	if !IsValidEmail(user.Email) {
		slog.Error("invalid email")
		api.WriteProblem(w, api.BadRequest("Email is not valid", ""))
		return
	}
	if !IsStrongPassword(user.Password) {
		slog.Error("password not strong enough")
		api.WriteProblem(w, api.BadRequest("Password must contain one uppercase character, one special character and one number", ""))
		return
	}

	err = s.RegisterUser(ctx, user)
	if err != nil{
		slog.Error("error creating the user","error",err.Error())
		api.WriteProblem(w, api.Internal("Error creating the user",""))
		return
	}

	slog.Info("user created successfully")

	resp:=&Response{
		Message: "User registered successfully",
	}

	api.WriteResponse(w, 201, resp)
}

func (s *Scope) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}
