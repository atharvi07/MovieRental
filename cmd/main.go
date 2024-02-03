package main

import (
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
	engine := gin.Default()
	route.RegisterRoute(dbConn, engine)
	err = engine.Run(appConfig.ServerAddress)
	if err != nil {
		log.Printf("Error starting server :  %v", err)
	}
}

//
//func populateDb() {
//	init.PopulateDB()
//}
