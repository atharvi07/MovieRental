package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie_rental/internal/app/service"
	"net/http"
)

type MovieHandler interface {
	GetMovies(ctx *gin.Context)
}

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return &movieHandler{
		movieService: movieService,
	}
}

func (movieHandler movieHandler) GetMovies(ctx *gin.Context) {

	genre := ctx.Query("genre")
	actor := ctx.Query("actor")
	year := ctx.Query("year")

	fmt.Println("genre ", genre, " actor ", actor, " year ", year)

	movies, err := movieHandler.movieService.GetMovies(genre, actor, year)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println("response ", movies)
	ctx.JSON(http.StatusOK, gin.H{
		"totalResults": len(movies),
		"data":         movies,
	})
}
