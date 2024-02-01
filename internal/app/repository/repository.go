package repository

import (
	"database/sql"
	"fmt"
	"movie_rental/internal/app/dto"
)

type MovieRepository interface {
	SaveMovies(movies []dto.MovieData)
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

	// dump to db
	fmt.Println("Movie Data Saved")
	/*
		insertStmt := fmt.Sprintf(`
				INSERT INTO %s (
					Title, Year, Rated, Released, Runtime, Genre, Director, Writer,
					Actors, Plot, Language, Country, Awards, Poster, MetaScore,
					IMDBRating, IMDBVotes, ImdbId, Type, DVD, BoxOffice, Production,
					Website, Response
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)
			`, tableName)

			_, err = db.Exec(
				insertStmt,
				movie.Title, movie.Year, movie.Rated, movie.Released, movie.Runtime,
				movie.Genre, movie.Director, movie.Writer, movie.Actors, movie.Plot,
				movie.Language, movie.Country, movie.Awards, movie.Poster, movie.MetaScore,
				movie.IMDBRating, movie.IMDBVotes, movie.ImdbId, movie.Type, movie.DVD,
				movie.BoxOffice, movie.Production, movie.Website, movie.Response,
			)
	*/

}
