package db

import (
	"database/sql"
	_ "github.com/lib/pq"
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
