package job

import (
	"time"

	"github.com/uptrace/bun"
)

type UserRegister struct {
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	Message string `json:"message"`
}

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64     `bun:",pk,autoincrement" json:"id"`
	Username      string    `bun:"username,notnull,unique" json:"username"`
	Email         string    `bun:"email,notnull,unique" json:"email"`
	Password      string    `bun:"password_hash,notnull" json:"-"`
	Role          string    `bun:"role,notnull,default:'user'" json:"role"`
	CreatedAt     time.Time `bun:"created_at,default:now()" json:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at,default:now()" json:"updated_at"`
}