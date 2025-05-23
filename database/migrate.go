package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func Migrate(db *sql.DB) {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create postgres driver: %v", err)
	}

	m, er := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if er != nil {
		log.Fatalf("migration init failed: %v", er)
	}

	if er = m.Up(); er != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", er)
	}

	log.Println("Migrations executed successfully")
}
