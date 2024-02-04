package route

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"movie_rental/internal/app/handler"
	"movie_rental/internal/app/repository"
	"movie_rental/internal/app/service"
	"net/http"
)

func RegisterRoute(dbConn *sql.DB, engine *gin.Engine) {

	movieRepo := repository.NewMovieRepo(dbConn)

	movieService := service.NewMovieService(movieRepo)

	movieHandler := handler.NewMovieHandler(movieService)

	engine.GET("/hello-world", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})

	group := engine.Group("/movieRental/api/v1")
	{
		group.GET("/movies", movieHandler.GetMovies)
		group.GET("/movie/:movieId", movieHandler.GetMovieById)
		group.GET("/add-to-cart/:movieId", movieHandler.AddToCart)
	}
}
