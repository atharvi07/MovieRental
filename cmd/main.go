package main

import (
	"github.com/gin-gonic/gin"
	"movie_rental/internal/app/route"
)

func main() {
	//fmt.Println("Hello World!!")
	engine := gin.Default()
	route.RegisterRoute(engine)
	engine.Run(":8080")
}
