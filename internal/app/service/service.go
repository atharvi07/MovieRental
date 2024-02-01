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
	/*
		https://www.omdbapi.com/?apikey=22b7636d&y=2024&s=Movie&page=1&type=Movie
		https://www.omdbapi.com/?apikey=22b7636d&y=2024&i=imbd&page=1&type=Movie
	*/
	var baseUrl = "https://www.omdbapi.com"
	var apiKey = "22b7636d"
	var pageNumber = "1"
	var searchUrl = baseUrl + "/?apikey=" + apiKey + "&s=Movie&type=Movie&y=2024" + "&page="

	searchMovieData := movieService.getSearchMovieData(searchUrl, pageNumber)

	totalResults, _ := strconv.Atoi(searchMovieData.TotalResults)
	var noOfPages = totalResults / 10 //34 -> 3 ... 1 ..3
	var remainder = totalResults % 10 // 4

	if remainder > 0 {
		noOfPages++ // 4 ... 1..4 // 4-1
	}

	//var wg sync.WaitGroup
	//wg.Add(noOfPages - 1)

	for i := 2; i <= noOfPages; i++ {
		pageNumber = strconv.Itoa(i)
		go movieService.getSearchMovieData(searchUrl, pageNumber)
	}
	//wg.Wait()

}

func (movieService movieService) getSearchMovieData(searchUrl string, pageNumber string) dto.SearchMovieData {

	//defer wg.Done()

	//fmt.Println(searchUrl)

	response, err := http.Get(searchUrl + pageNumber)

	if err != nil {
		return dto.SearchMovieData{}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return dto.SearchMovieData{}
	}

	//fmt.Println("Response status: ", response.StatusCode)
	//fmt.Println("Response body: ", string(body))

	var searchMovieData dto.SearchMovieData
	err = json.Unmarshal(body, &searchMovieData)

	if err != nil {
		fmt.Println("Parsing error : ", err)
		log.Fatal(err)
	}
	//fmt.Println("SearchMovieData :", searchMovieData)

	movieService.fetchMovieDataFromOmdbAndSaveData(searchMovieData)

	return searchMovieData
}

func (movieService movieService) fetchMovieDataFromOmdbAndSaveData(searchMovieData dto.SearchMovieData) {

	var wg sync.WaitGroup
	wg.Add(len(searchMovieData.MovieDetails))

	dataChannel := make(chan dto.MovieData)
	var movies []dto.MovieData
	for _, detail := range searchMovieData.MovieDetails {
		//fmt.Println(detail.ImdbID)
		go movieService.getMovieDetailsFromOmdbService(detail.ImdbID, &wg, dataChannel)
		movies = append(movies, <-dataChannel)
	}
	wg.Wait()
	close(dataChannel)

	fmt.Println("Movies : ", movies)

	// savebulk for every page
	movieService.repository.SaveMovies(movies)
}

func (movieService movieService) getMovieDetailsFromOmdbService(imdbId string, wg *sync.WaitGroup, dataChannel chan dto.MovieData) {
	/*
		https://www.omdbapi.com/?apikey=22b7636d&i=tt1674782
	*/
	defer wg.Done()

	var baseUrl = "https://www.omdbapi.com"
	var apiKey = "22b7636d"

	var movieDetailUrl = baseUrl + "/?apikey=" + apiKey + "&i=" + imdbId
	//fmt.Println(movieDetailUrl)
	response, err := http.Get(movieDetailUrl)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	//fmt.Println("Response status: ", response.StatusCode)
	//fmt.Println("Response body: ", string(body))

	var movieData dto.MovieData
	err = json.Unmarshal(body, &movieData)
	if err != nil {
		fmt.Println("Parsing error in get movie details")
		log.Fatal(err)
	}

	dataChannel <- movieData

}
