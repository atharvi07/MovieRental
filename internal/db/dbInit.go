package db

import (
	"database/sql"
	"fmt"
	"log"
)

type DbInit interface {
	IntializeDb()
}

type dbInit struct {
	db *sql.DB
}

func NewDbInit(db *sql.DB) DbInit {
	return &dbInit{
		db: db,
	}
}

func (dbInit dbInit) IntializeDb() {
	createTableSQL := `
CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    year VARCHAR(4),
    rated VARCHAR(20),
    released VARCHAR(20),
    runtime VARCHAR(20),
    genre VARCHAR(255),
    director VARCHAR(255),
    writer VARCHAR(255),
    actors VARCHAR(255),
    plot TEXT,
    language VARCHAR(255),
    country VARCHAR(255),
    awards VARCHAR(255),
    poster VARCHAR(255),
    ratings JSONB,
    metascore VARCHAR(10),
    imdb_rating VARCHAR(10),
    imdb_votes VARCHAR(20),
    imdb_id VARCHAR(20),
    movie_type VARCHAR(20),
    dvd VARCHAR(20),
    box_office VARCHAR(50),
    production VARCHAR(255),
    website VARCHAR(255),
    response VARCHAR(10)
);
`

	result, err := dbInit.db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting RowsAffected:", err)
		return
	}

	fmt.Printf("RowsAffected: %d\n", rowsAffected)
}
