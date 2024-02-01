package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func CreateConnection() *sql.DB {
	dataSource := fmt.Sprintf(
		"postgres://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"root",
	)
	fmt.Println("datasource : ", dataSource)
	dbConn, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal("unable to open connection with database ", err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping database ", err.Error())
	}
	return dbConn
}
