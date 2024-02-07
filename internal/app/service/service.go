package service

import (
	"movie_rental/internal/app/dto"
	"movie_rental/internal/app/repository"
)

type MovieService interface {
	GetMovies(genre string, actor string, year string) ([]dto.MovieData, error)
	GetMovieById(id string) (*dto.MovieDetails, error)
	AddToCart(movieID string) error
	GetCart() dto.Cart
}

type movieService struct {
	repository repository.MovieRepository
	cart       dto.Cart
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{
		repository: movieRepository,
	}
}

func (movieService *movieService) GetMovies(genre string, actor string, year string) ([]dto.MovieData, error) {
	movies, err := movieService.repository.FindMovies(genre, actor, year)
	if err != nil {
		return []dto.MovieData{}, err
	}

	return movies, nil
}

func (movieService *movieService) GetMovieById(movieId string) (*dto.MovieDetails, error) {

	movie, err := movieService.repository.FindMovieById(movieId)
	if err != nil {

		return &dto.MovieDetails{}, err
	}
	return movie, nil
}

func (movieService *movieService) AddToCart(movieID string) error {
	movie, err := movieService.repository.FindMovieById(movieID)
	if err != nil {
		return err
	}

	movieService.cart.Movies = append(movieService.cart.Movies, *movie)
	return nil
}

func (movieService *movieService) GetCart() dto.Cart {
	return movieService.cart
}
