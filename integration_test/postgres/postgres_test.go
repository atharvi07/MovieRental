package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"movie_rental/internal/app/dto"
	"movie_rental/internal/app/route"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testDbInstance *sql.DB

func TestMain(m *testing.M) {
	testDB := SetupTestDatabase("postgres://localhost:5432/postgres?user=postgres&password=root&sslmode=disable")
	testDbInstance = testDB.db
	defer testDB.TearDown()
	os.Exit(m.Run())
}

func TestFindMovies(t *testing.T) {
	t.Run("should return list of movies when find movies call is successful", func(t *testing.T) {
		engine := gin.Default()
		route.RegisterRoute(testDbInstance, engine)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movies", nil)
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

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.Equal(t, response.TotalResults, 13)
	})

	t.Run("should return empty list of movies when no movies available", func(t *testing.T) {
		engine := gin.Default()
		route.RegisterRoute(testDbInstance, engine)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movies?year=2023", nil)
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

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.Equal(t, response.TotalResults, 0)
	})
}

func TestFindMovieById(t *testing.T) {
	t.Run("should return a movies when find movie by id call is successful", func(t *testing.T) {
		engine := gin.Default()
		route.RegisterRoute(testDbInstance, engine)

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
			Data dto.MovieData `json:"data"`
		}
		if err = json.Unmarshal(body, &response); err != nil {
			fmt.Println("Unable to parse json")
		}

		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, response.Data.ImdbId, "tt5779228")
	})

	t.Run("should return error when invalid movie id is passed", func(t *testing.T) {
		engine := gin.Default()
		route.RegisterRoute(testDbInstance, engine)

		request, err := http.NewRequest(http.MethodGet, "/movieRental/api/v1/movie/1234qewe", nil)
		require.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	})
}
