package postgres

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"movie_rental/internal/app/repository"
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
		repo := repository.NewMovieRepo(testDbInstance)

		movies, err := repo.FindMovies("", "", "")

		assert.NoError(t, err)
		assert.NotNil(t, movies)
	})
	t.Run("should return empty list of movies when no movies available", func(t *testing.T) {
		repo := repository.NewMovieRepo(testDbInstance)

		movies, err := repo.FindMovies("", "", "2023")

		assert.NoError(t, err)
		assert.Nil(t, movies)
	})
}

func TestFindMovieById(t *testing.T) {
	t.Run("should return a movies when find movie by id call is successful", func(t *testing.T) {
		repo := repository.NewMovieRepo(testDbInstance)

		movie, err := repo.FindMovieById("tt5779228")

		assert.NoError(t, err)
		assert.NotNil(t, movie)
	})
	t.Run("should return error when invalid movie id is passed", func(t *testing.T) {
		repo := repository.NewMovieRepo(testDbInstance)

		_, err := repo.FindMovieById("1234qwer")

		assert.Error(t, err)
	})
}
