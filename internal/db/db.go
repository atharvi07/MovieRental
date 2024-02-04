package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

func CreateConnection(dbDriver string, dbSource string) *sql.DB {
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("unable to open connection with database ", err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping database ", err.Error())
	}
	return dbConn
}

func ApplyMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println("Driver not installed", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	if err != nil {
		return fmt.Errorf("unable to create migration: %v", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to apply migrations %v", err)
	}
	return nil
}
