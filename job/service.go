package job

import (
	"context"
	"log/slog"
)

func (s *Scope) RegisterUser(ctx context.Context, user UserRegister) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		slog.Error("error hashing the password")
		return err
	}

	err = s.db.CreateUser(ctx, user.UserName, user.Email, hashedPassword)
	if err != nil{
		slog.Error("error inserting the user")
		return err
	}

	return nil
}