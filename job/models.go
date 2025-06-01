package job

type UserRegister struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json: status`
	Details  string `json:"details,omitempty"`
	Instance string `json:"instance,omitempty"`
}