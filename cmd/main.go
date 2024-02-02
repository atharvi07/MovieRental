package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"movie_rental/internal/app/route"
)

func main() {
	engine := gin.Default()
	route.RegisterRoute(engine)
	err := engine.Run(":8080")
	if err != nil {
		log.Printf("Error starting server :  %v", err)
	}

}

//
//func populateDb() {
//	init.PopulateDB()
//}
