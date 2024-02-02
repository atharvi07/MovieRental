package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"movie_rental/internal/app/dto"
	"movie_rental/internal/app/repository/mocks"
	"testing"
)

func TestMovieService_GetMovies(t *testing.T) {

	t.Run("should return movies on successful of fetch data for movies", func(t *testing.T) {
		mockRepository := &mocks.MovieRepository{}
		movie2 := dto.MovieData{
			Id:         "1",
			Title:      "Untitled Karate Kid Movie",
			Year:       "2024",
			Rated:      "",
			Released:   "13 Dec 2024",
			Runtime:    "",
			Genre:      "Action, Drama, Family",
			Director:   "Jonathan Entwistle",
			Writer:     "Rob Lieber",
			Actors:     "Jackie Chan, Ralph Macchio",
			Plot:       "Plot under wraps.",
			Language:   "English",
			Country:    "United States",
			Awards:     "N/A",
			Poster:     "N/A",
			Ratings:    nil,
			MetaScore:  "",
			IMDBRating: "N/A",
			IMDBVotes:  "",
			ImdbId:     "tt1674782",
			Type:       "movie",
			DVD:        "",
			BoxOffice:  "",
			Production: "",
			Website:    "",
			Response:   "",
		}
		movie1 := dto.MovieData{
			Id:         "2",
			Title:      "The Garfield Movie",
			Year:       "2024",
			Rated:      "",
			Released:   "24 May 2024",
			Runtime:    "",
			Genre:      "Animation, Adventure, Comedy",
			Director:   "Mark Dindal",
			Writer:     "Paul A. Kaplan, Mark Torgove, David Reynolds",
			Actors:     "Hannah Waddingham, Chris Pratt, Nicholas Hoult",
			Plot:       "Garfield is about to go on a wild outdoor adventure. After an unexpected reunion with his long-lost father - the cat Vic - Garfield and Odie are forced to abandon their pampered life to join Vic in a hilarious, high-stakes heist.",
			Language:   "English",
			Country:    "United Kingdom, United States",
			Awards:     "N/A",
			Poster:     "https://m.media-amazon.com/images/M/MV5BNzk4ODdiOTEtMTk3YS00MzZmLTgyOWMtYzc1NjgxYWE2MmMyXkEyXkFqcGdeQXVyMTUzMTg2ODkz._V1_SX300.jpg",
			Ratings:    nil,
			MetaScore:  "",
			IMDBRating: "N/A",
			IMDBVotes:  "",
			ImdbId:     "tt5779228",
			Type:       "movie",
			DVD:        "",
			BoxOffice:  "",
			Production: "",
			Website:    "",
			Response:   "",
		}
		mockRepository.On("FindMovies", "", "", "").Return([]dto.MovieData{movie1, movie2}, nil)

		movieService := NewMovieService(mockRepository)

		movies, err := movieService.GetMovies("", "", "")

		assert.NoError(t, err)
		assert.Equal(t, []dto.MovieData{movie1, movie2}, movies)
	})

	t.Run("should return error when unable to fetch data for movies", func(t *testing.T) {
		mockRepository := &mocks.MovieRepository{}
		mockRepository.On("FindMovies", "", "", "").Return(nil, errors.New("simulated internal server error"))

		movieService := NewMovieService(mockRepository)

		movies, err := movieService.GetMovies("", "", "")

		assert.Error(t, err)
		assert.Equal(t, []dto.MovieData{}, movies)
	})
}
