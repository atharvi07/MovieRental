package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie_rental/internal/app/service"
	"net/http"
)

type MovieHandler interface {
	GetMovies(ctx *gin.Context)
	GetMovieById(ctx *gin.Context)
	AddToCart(ctx *gin.Context)
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
	ctx.JSON(http.StatusOK, gin.H{
		"totalResults": len(movies),
		"data":         movies,
	})
}

func (movieHandler movieHandler) GetMovieById(ctx *gin.Context) {
	movieId := ctx.Param("movieId")

	movie, err := movieHandler.movieService.GetMovieById(movieId)
	if err != nil {
		if err.Error() == "invalid id passed" {
			ctx.String(http.StatusBadRequest, err.Error())
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": movie,
	})

}

func (movieHandler movieHandler) AddToCart(ctx *gin.Context) {
	movieId := ctx.Param("movieId")
	err := movieHandler.movieService.AddToCart(movieId)
	if err != nil {
		if err.Error() == "invalid id passed" {
			fmt.Println("Inside invalid id error")
			ctx.String(http.StatusBadRequest, err.Error())
		}
		fmt.Println("Inside internal server error")
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, movieHandler.movieService.GetCart())
}
