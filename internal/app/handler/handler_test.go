package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"movie_rental/internal/app/dto"
	"movie_rental/internal/app/repository/mocks"
	"movie_rental/internal/app/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovieHandler_GetMovies(t *testing.T) {
	t.Run("should return all movies when on find movies successful call", func(t *testing.T) {
		engine := gin.Default()
		mockRepository := &mocks.MovieRepository{}
		movie2 := dto.MovieData{
			Id:     "1",
			Title:  "Untitled Karate Kid Movie",
			Year:   "2024",
			Genre:  "Action, Drama, Family",
			Actors: "Jackie Chan, Ralph Macchio",
			Plot:   "Plot under wraps.",
			Poster: "N/A",
			ImdbId: "tt1674782",
		}
		movie1 := dto.MovieData{
			Id:     "2",
			Title:  "The Garfield Movie",
			Year:   "2024",
			Genre:  "Animation, Adventure, Comedy",
			Actors: "Hannah Waddingham, Chris Pratt, Nicholas Hoult",
			Plot:   "Garfield is about to go on a wild outdoor adventure. After an unexpected reunion with his long-lost father - the cat Vic - Garfield and Odie are forced to abandon their pampered life to join Vic in a hilarious, high-stakes heist.",
			Poster: "https://m.media-amazon.com/images/M/MV5BNzk4ODdiOTEtMTk3YS00MzZmLTgyOWMtYzc1NjgxYWE2MmMyXkEyXkFqcGdeQXVyMTUzMTg2ODkz._V1_SX300.jpg",
			ImdbId: "tt5779228",
		}
		mockRepository.On("FindMovies", "", "", "").Return([]dto.MovieData{movie1, movie2}, nil)

		movieService := service.NewMovieService(mockRepository)
		handler := NewMovieHandler(movieService)

		engine.GET("/movieRental/api/v1/movies", handler.GetMovies)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movies", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		assert.Equal(t, responseRecorder.Code, http.StatusOK)
	})

	t.Run("should return filtered movies when genre is passed", func(t *testing.T) {
		engine := gin.Default()
		mockRepository := &mocks.MovieRepository{}
		movie1 := dto.MovieData{
			Id:     "2",
			Title:  "The Garfield Movie",
			Year:   "2024",
			Genre:  "Animation, Adventure, Comedy",
			Actors: "Hannah Waddingham, Chris Pratt, Nicholas Hoult",
			Plot:   "Garfield is about to go on a wild outdoor adventure. After an unexpected reunion with his long-lost father - the cat Vic - Garfield and Odie are forced to abandon their pampered life to join Vic in a hilarious, high-stakes heist.",
			Poster: "https://m.media-amazon.com/images/M/MV5BNzk4ODdiOTEtMTk3YS00MzZmLTgyOWMtYzc1NjgxYWE2MmMyXkEyXkFqcGdeQXVyMTUzMTg2ODkz._V1_SX300.jpg",
			ImdbId: "tt5779228",
		}
		movie2 := dto.MovieData{
			Id:     "1",
			Title:  "Untitled Karate Kid Movie",
			Year:   "2024",
			Genre:  "Action, Drama, Family",
			Actors: "Jackie Chan, Ralph Macchio",
			Plot:   "Plot under wraps.",
			Poster: "N/A",
			ImdbId: "tt1674782",
		}
		mockRepository.On("FindMovies", "Action", "", "").Return([]dto.MovieData{movie1, movie2}, nil)

		movieService := service.NewMovieService(mockRepository)
		handler := NewMovieHandler(movieService)

		engine.GET("/movieRental/api/v1/movies", handler.GetMovies)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movies?genre=Action", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		body, err := ioutil.ReadAll(responseRecorder.Body)
		if err != nil {
			fmt.Println("Error reading responseRecorder body:", err)
			return
		}
		var response struct {
			TotalResults int             `json:"totalResults"`
			Data         []dto.MovieData `json:"data"`
		}
		if err = json.Unmarshal(body, &response); err != nil {
			fmt.Println("Unable to parse json")
		}
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, response.Data[1].Title, "Untitled Karate Kid Movie")
	})

	t.Run("should return filtered movies when actor is passed", func(t *testing.T) {
		engine := gin.Default()
		mockRepository := &mocks.MovieRepository{}
		movie1 := dto.MovieData{
			Id:     "2",
			Title:  "The Garfield Movie",
			Year:   "2024",
			Genre:  "Animation, Adventure, Comedy",
			Actors: "Hannah Waddingham, Chris Pratt, Nicholas Hoult",
			Plot:   "Garfield is about to go on a wild outdoor adventure. After an unexpected reunion with his long-lost father - the cat Vic - Garfield and Odie are forced to abandon their pampered life to join Vic in a hilarious, high-stakes heist.",
			Poster: "https://m.media-amazon.com/images/M/MV5BNzk4ODdiOTEtMTk3YS00MzZmLTgyOWMtYzc1NjgxYWE2MmMyXkEyXkFqcGdeQXVyMTUzMTg2ODkz._V1_SX300.jpg",
			ImdbId: "tt5779228",
		}
		movie2 := dto.MovieData{
			Id:     "1",
			Title:  "Untitled Karate Kid Movie",
			Year:   "2024",
			Genre:  "Action, Drama, Family",
			Actors: "Jackie Chan, Ralph Macchio",
			Plot:   "Plot under wraps.",
			Poster: "N/A",
			ImdbId: "tt1674782",
		}
		mockRepository.On("FindMovies", "", "Hannah Waddingham", "").Return([]dto.MovieData{movie1, movie2}, nil)

		movieService := service.NewMovieService(mockRepository)
		handler := NewMovieHandler(movieService)

		engine.GET("/movieRental/api/v1/movies", handler.GetMovies)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movies?actor=Hannah Waddingham", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		body, err := ioutil.ReadAll(responseRecorder.Body)
		if err != nil {
			fmt.Println("Error reading responseRecorder body:", err)
			return
		}
		var response struct {
			TotalResults int             `json:"totalResults"`
			Data         []dto.MovieData `json:"data"`
		}

		if err = json.Unmarshal(body, &response); err != nil {
			fmt.Println("Unable to parse json")
		}
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, response.Data[0].Title, "The Garfield Movie")
	})

	t.Run("should return internal server error when unable to fetch movies from database", func(t *testing.T) {
		engine := gin.Default()
		mockRepository := &mocks.MovieRepository{}
		mockRepository.On("FindMovies", "", "", "").Return(nil, errors.New("simulated internal server error"))

		movieService := service.NewMovieService(mockRepository)
		handler := NewMovieHandler(movieService)

		engine.GET("/movieRental/api/v1/movies", handler.GetMovies)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movies", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		assert.Equal(t, responseRecorder.Code, http.StatusInternalServerError)
	})
}

func TestMovieHandler_GetMovieById(t *testing.T) {
	t.Run("should return a movie on success of get movie by id ", func(t *testing.T) {
		engine := gin.Default()
		mockRepository := &mocks.MovieRepository{}
		movie1 := &dto.MovieDetails{
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
		mockRepository.On("FindMovieById", "tt5779228").Return(movie1, nil)

		movieService := service.NewMovieService(mockRepository)
		handler := NewMovieHandler(movieService)

		engine.GET("/movieRental/api/v1/movie/:movieId", handler.GetMovieById)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movie/tt5779228", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)
		body, err := ioutil.ReadAll(responseRecorder.Body)
		if err != nil {
			fmt.Println("Error reading responseRecorder body:", err)
			return
		}
		var response struct {
			Data dto.MovieDetails `json:"data"`
		}
		if err = json.Unmarshal(body, &response); err != nil {
			fmt.Println("Unable to parse json")
		}

		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, response.Data.Title, movie1.Title)
	})

	t.Run("should return bad request error when invalid id passed", func(t *testing.T) {
		engine := gin.Default()
		mockRepository := &mocks.MovieRepository{}
		mockRepository.On("FindMovieById", "tt57228").Return(&dto.MovieDetails{}, errors.New("invalid id passed"))

		movieService := service.NewMovieService(mockRepository)
		handler := NewMovieHandler(movieService)

		engine.GET("/movieRental/api/v1/movie/:movieId", handler.GetMovieById)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movie/tt57228", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	})
}

func TestMovieHandler_AddToCart(t *testing.T) {
	t.Run("should add a movie to a cart for valid id ", func(t *testing.T) {
		engine := gin.Default()
		mockRepository := &mocks.MovieRepository{}
		movie1 := &dto.MovieDetails{
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
		mockRepository.On("FindMovieById", "tt5779228").Return(movie1, nil)

		movieService := service.NewMovieService(mockRepository)
		handler := NewMovieHandler(movieService)

		engine.GET("/movieRental/api/v1/add-to-cart/:movieId", handler.AddToCart)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/add-to-cart/tt5779228", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)
		body, err := ioutil.ReadAll(responseRecorder.Body)
		if err != nil {
			fmt.Println("Error reading responseRecorder body:", err)
			return
		}
		var response struct {
			Data []dto.MovieDetails `json:"movies"`
		}
		if err = json.Unmarshal(body, &response); err != nil {
			fmt.Println("Unable to parse json")
		}

		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, response.Data[0].Title, movie1.Title)
	})

	t.Run("should return bad request error when invalid id passed", func(t *testing.T) {
		engine := gin.Default()
		mockRepository := &mocks.MovieRepository{}
		mockRepository.On("FindMovieById", "13wdewe").Return(&dto.MovieDetails{}, errors.New("invalid id passed"))

		movieService := service.NewMovieService(mockRepository)
		handler := NewMovieHandler(movieService)

		engine.GET("/movieRental/api/v1/add-to-cart/:movieId", handler.AddToCart)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/add-to-cart/13wdewe", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	})
}
