package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie_rental/internal/app/handler"
	"movie_rental/internal/app/repository"
	"movie_rental/internal/app/service"
	"movie_rental/internal/db"
	"net/http"
)

func RegisterRoute(engine *gin.Engine) {

	dbConn := db.CreateConnection()
	dbInit := db.NewDbInit(dbConn)
	dbInit.IntializeDb()

	movieRepo := repository.NewMovieRepo(dbConn)

	movieService := service.NewMovieService(movieRepo)
	movieService.PopulateDatabase()

	handler.NewMovieHandler(movieService)

	engine.GET("/hello-world", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})
	fmt.Println("Application Running on 8080")
}
