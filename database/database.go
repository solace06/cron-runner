package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/solace06/cron-runner/job/config"
)

type DB struct {
	Conn *bun.DB
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

	db := pgdriver.NewConnector(pgdriver.WithDSN(connStr))

	dbConn := bun.NewDB(
		sql.OpenDB(db),
		pgdialect.New(),
	)

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
