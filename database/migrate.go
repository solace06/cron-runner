package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func (db *DB) Migrate() {

	driver, err := postgres.WithInstance(db.Conn, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create postgres driver: %v", err)
	}

	m, er := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", driver)
	if er != nil {
		log.Fatalf("migration init failed: %v", er)
	}

	if er = m.Up(); er != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", er)
	}

	log.Println("Migrations executed successfully")
}
