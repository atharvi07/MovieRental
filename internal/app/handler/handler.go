package handler

import "movie_rental/internal/app/service"

type MovieHandler interface {
}

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return &movieHandler{
		movieService: movieService,
	}
}
