package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"movie_rental/config"
	"movie_rental/internal/app/route"
	"movie_rental/internal/db"
)

func main() {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load app configurations")
	}
	dbConn := db.CreateConnection(appConfig.DBDriver, appConfig.DBSource)

	err = db.ApplyMigrations(dbConn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrations applied")

	engine := gin.Default()
	route.RegisterRoute(dbConn, engine)
	err = engine.Run(appConfig.ServerAddress)
	if err != nil {
		log.Printf("Error starting server :  %v", err)
	}
}
