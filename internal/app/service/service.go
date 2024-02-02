package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"movie_rental/internal/app/dto"
	"movie_rental/internal/app/repository"
	"net/http"
	"strconv"
	"sync"
)

type MovieService interface {
	PopulateDatabase()
	GetMovies(genre string, actor string, year string) ([]dto.MovieData, error)
}

type movieService struct {
	repository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{
		movieRepository,
	}
}

func (movieService movieService) PopulateDatabase() {

	movieService.fetchAndSaveMovieData()

}

func (movieService movieService) fetchAndSaveMovieData() {
	var baseUrl = "https://www.omdbapi.com"
	var apiKey = "22b7636d"
	var pageNumber = "1"
	var searchUrl = baseUrl + "/?apikey=" + apiKey + "&s=Movie&type=Movie&y=2024" + "&page="

	searchMovieData := movieService.getSearchMovieData(searchUrl, pageNumber)

	totalResults, _ := strconv.Atoi(searchMovieData.TotalResults)
	var noOfPages = totalResults / 10
	var remainder = totalResults % 10

	if remainder > 0 {
		noOfPages++
	}

	for i := 2; i <= noOfPages; i++ {
		pageNumber = strconv.Itoa(i)
		go movieService.getSearchMovieData(searchUrl, pageNumber)
	}

}

func (movieService movieService) getSearchMovieData(searchUrl string, pageNumber string) dto.SearchMovieData {

	response, err := http.Get(searchUrl + pageNumber)

	if err != nil {
		return dto.SearchMovieData{}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return dto.SearchMovieData{}
	}

	var searchMovieData dto.SearchMovieData
	err = json.Unmarshal(body, &searchMovieData)

	if err != nil {
		fmt.Println("Parsing error : ", err)
		log.Fatal(err)
	}

	movieService.fetchMovieDataFromOmdbAndSaveData(searchMovieData)

	return searchMovieData
}

func (movieService movieService) fetchMovieDataFromOmdbAndSaveData(searchMovieData dto.SearchMovieData) {

	var wg sync.WaitGroup
	wg.Add(len(searchMovieData.MovieDetails))

	dataChannel := make(chan dto.MovieData)
	var movies []dto.MovieData
	for _, detail := range searchMovieData.MovieDetails {
		go movieService.getMovieDetailsFromOmdbService(detail.ImdbID, &wg, dataChannel)
		movies = append(movies, <-dataChannel)
	}
	wg.Wait()
	close(dataChannel)

	fmt.Println("Movies : ", movies)

	movieService.repository.SaveMovies(movies)
}

func (movieService movieService) getMovieDetailsFromOmdbService(imdbId string, wg *sync.WaitGroup, dataChannel chan dto.MovieData) {

	defer wg.Done()

	var baseUrl = "https://www.omdbapi.com"
	var apiKey = "22b7636d"

	var movieDetailUrl = baseUrl + "/?apikey=" + apiKey + "&i=" + imdbId

	response, err := http.Get(movieDetailUrl)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var movieData dto.MovieData
	err = json.Unmarshal(body, &movieData)
	if err != nil {
		fmt.Println("Parsing error in get movie details")
		log.Fatal(err)
	}

	dataChannel <- movieData

}

func (movieService movieService) GetMovies(genre string, actor string, year string) ([]dto.MovieData, error) {
	movies, err := movieService.repository.FindMovies(genre, actor, year)
	if err != nil {
		return []dto.MovieData{}, err
	}

	return movies, nil
}
