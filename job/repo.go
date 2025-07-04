package job

import (
	"context"
	"log/slog"
)

func (s *Scope) CreateUser(ctx context.Context, username string, email string, password string) error {
	user := &User{
		Username: username,
		Email:    email,
		Password: password,
	}

	_, err := s.db.Conn.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		slog.Error("error creating user", "err", err.Error())
		return err
	}
	return nil
}