package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"
)

type TestDatabase struct {
	db        *sql.DB
	container testcontainers.Container
}

func SetupTestDatabase(dbSource string) *TestDatabase {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	container, db, err := createPostgresContainer(ctx, dbSource)
	if err != nil {
		log.Fatal("failed to setup test ", err)
	}

	err = migrateDb(db)
	if err != nil {
		log.Fatal("failed to perform db migration", err)
	}
	cancel()

	return &TestDatabase{
		container: container,
		db:        db,
	}
}

func createPostgresContainer(ctx context.Context, dbSource string) (testcontainers.Container, *sql.DB, error) {
	request := testcontainers.GenericContainerRequest{ContainerRequest: testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "postgres",
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "root",
		},
		WaitingFor: wait.ForSQL("5432", "postgres", func(host string, port nat.Port) string {
			return dbSource
		}).WithStartupTimeout(5 * time.Second),
		Name: "postgres",
	}, Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, request)
	if err != nil {
		return container, nil, fmt.Errorf("failed to start container: %v", err)
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return container, nil, fmt.Errorf("failed to get container external port: %v", err)
	}
	log.Println("postgres container ready and running at port: ", mappedPort.Port())
	dbConn, err := sql.Open("postgres", dbSource)
	if err != nil {
		return container, dbConn, fmt.Errorf("failed to establish database connection: %v", err)
	}
	return container, dbConn, nil
}

func migrateDb(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file:///Users/atharvimali/Documents/learning/Golang/movie_rental/internal/db/migrations", "postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("migration done")

	return nil
}

func (tdb *TestDatabase) TearDown() {
	tdb.db.Close()
	_ = tdb.container.Terminate(context.Background())
}
