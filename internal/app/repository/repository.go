package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"movie_rental/internal/app/dto"
	"strconv"
	"strings"
)

type MovieRepository interface {
	SaveMovies(movies []dto.MovieData)
	FindMovies(genre string, actor string, year string) ([]dto.MovieData, error)
	FindMovieById(id string) (dto.MovieData, error)
}

type movieRepository struct {
	*sql.DB
}

func NewMovieRepo(db *sql.DB) MovieRepository {
	return &movieRepository{
		db,
	}
}

func (movieRepository movieRepository) SaveMovies(movies []dto.MovieData) {

	var tableName = "movies"

	insertStmt := fmt.Sprintf(`
				INSERT INTO %s (
					 title, year, rated, released, runtime, genre, director, writer, 
		actors, plot, language, country, awards, poster, metascore, 
		imdb_rating, imdb_votes, imdb_id, movie_type, dvd, box_office, 
		production, website, response
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)
			`, tableName)

	fmt.Println(insertStmt)

	for _, movie := range movies {
		_, err := movieRepository.DB.Exec(
			insertStmt,
			movie.Title, movie.Year, movie.Rated, movie.Released, movie.Runtime,
			movie.Genre, movie.Director, movie.Writer, movie.Actors, movie.Plot,
			movie.Language, movie.Country, movie.Awards, movie.Poster, movie.MetaScore,
			movie.IMDBRating, movie.IMDBVotes, movie.ImdbId, movie.Type, movie.DVD,
			movie.BoxOffice, movie.Production, movie.Website, movie.Response,
		)

		if err != nil {
			// If an error occurs, trigger a panic to rollback the transaction
			panic(err)
		}
	}

}

func (movieRepository movieRepository) FindMovies(genre string, actor string, year string) ([]dto.MovieData, error) {
	query := "SELECT id, title, year, released, genre, director, writer, actors, plot, language, country, awards, poster, imdb_rating, imdb_id, movie_type FROM movies m"

	conditions := []string{}
	args := []interface{}{}
	paramCount := 1

	if genre != "" {
		conditions = append(conditions, "m.genre LIKE $"+strconv.Itoa(paramCount))
		args = append(args, "%"+genre+"%")
		paramCount++
	}

	if actor != "" {
		conditions = append(conditions, "m.actors LIKE $"+strconv.Itoa(paramCount))
		args = append(args, "%"+actor+"%")
		paramCount++
	}

	if year != "" {
		conditions = append(conditions, "m.year = $"+strconv.Itoa(paramCount))
		args = append(args, year)
		paramCount++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	fmt.Println("Query:", query)

	rows, err := movieRepository.DB.Query(query, args...)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return nil, err
	}
	defer rows.Close()

	var movies []dto.MovieData

	for rows.Next() {
		var movie dto.MovieData
		err := rows.Scan(
			&movie.Id,
			&movie.Title, &movie.Year, &movie.Released,
			&movie.Genre, &movie.Director, &movie.Writer, &movie.Actors, &movie.Plot,
			&movie.Language, &movie.Country, &movie.Awards, &movie.Poster,
			&movie.IMDBRating, &movie.ImdbId,
			&movie.Type,
		)
		if err != nil {
			fmt.Println("Error occurred while row scan:", err)
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (movieRepository movieRepository) FindMovieById(id string) (dto.MovieData, error) {

	row := movieRepository.DB.QueryRow("SELECT id, title, year, released, genre, director, writer, actors, plot, language, country, awards, poster,imdb_rating, imdb_id, movie_type FROM movies m WHERE m.imdb_id = $1", id)

	var movie dto.MovieData
	err := row.Scan(
		&movie.Id,
		&movie.Title, &movie.Year, &movie.Released,
		&movie.Genre, &movie.Director, &movie.Writer, &movie.Actors, &movie.Plot,
		&movie.Language, &movie.Country, &movie.Awards, &movie.Poster,
		&movie.IMDBRating, &movie.ImdbId,
		&movie.Type,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.MovieData{}, errors.New("invalid id passed")
		}

		fmt.Println("error occurred while row scan : ", err)
		return dto.MovieData{}, err
	}

	return movie, nil

}
