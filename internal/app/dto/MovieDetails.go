package dto

type MovieDetails struct {
	Id         string
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	MetaScore  string   `json:"Metascore"`
	IMDBRating string   `json:"imdbRating"`
	IMDBVotes  string   `json:"imdbVotes"`
	ImdbId     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type MovieData struct {
	Id     string
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Genre  string `json:"Genre"`
	Actors string `json:"Actors"`
	Plot   string `json:"Plot"`
	Poster string `json:"Poster"`
	ImdbId string `json:"imdbID"`
}