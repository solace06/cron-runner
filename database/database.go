package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/solace06/cron-runner/job/config"
)

type DB struct {
	Conn    *bun.DB
	SQLConn *sql.DB
}

func NewDB(cfg *config.Config) (*DB, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	db := pgdriver.NewConnector(pgdriver.WithDSN(connStr))

	sqldb := sql.OpenDB(db)

	dbConn := bun.NewDB(
		sqldb,
		pgdialect.New(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := dbConn.PingContext(ctx); err != nil {
		return nil, err
	}

	slog.Info("Successfully connected to the database")

	return &DB{
		Conn:    dbConn,
		SQLConn: sqldb,
	}, nil
}

func (db *DB) Close() error {
	return db.Conn.Close()
}