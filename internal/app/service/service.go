package service

import (
	"movie_rental/internal/app/dto"
	"movie_rental/internal/app/repository"
)

type MovieService interface {
	GetMovies(genre string, actor string, year string) ([]dto.MovieData, error)
	GetMovieById(id string) (dto.MovieData, error)
}

type movieService struct {
	repository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{
		movieRepository,
	}
}

func (movieService movieService) GetMovies(genre string, actor string, year string) ([]dto.MovieData, error) {
	movies, err := movieService.repository.FindMovies(genre, actor, year)
	if err != nil {
		return []dto.MovieData{}, err
	}

	return movies, nil
}

func (movieService movieService) GetMovieById(movieId string) (dto.MovieData, error) {

	movie, err := movieService.repository.FindMovieById(movieId)
	if err != nil {

		return dto.MovieData{}, err
	}
	return movie, nil
}
