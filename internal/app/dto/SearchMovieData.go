package dto

type SearchMovieData struct {
	MovieDetails []MovieDetails `json:"Search"`
	TotalResults string         `json:"totalResults"`
}

type MovieDetails struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
