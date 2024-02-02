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

	movieRepo := repository.NewMovieRepo(dbConn)

	movieService := service.NewMovieService(movieRepo)
	//movieService.PopulateDatabase()

	movieHandler := handler.NewMovieHandler(movieService)

	engine.GET("/hello-world", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})

	group := engine.Group("/movieRental/api/v1")
	{
		group.GET("/movies", movieHandler.GetMovies)
		group.GET("/movie/:movieId", movieHandler.GetMovieById)

	}

	fmt.Println("Application Running on 8080")
}
