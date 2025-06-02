package job

type UserRegister struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}