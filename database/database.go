package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"
	_ "github.com/lib/pq" 

	"github.com/solace06/cron-runner/job/config"
)

type DB struct {
	Conn *sql.DB
}

func NewDB(cfg *config.Config) (*DB, error) {

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := dbConn.PingContext(ctx); err != nil {
		return nil, err
	}

	slog.Info("Successfully connected to the database")

	return &DB{Conn: dbConn}, nil
}

func (db *DB) Close() error {
	return db.Conn.Close()
}
