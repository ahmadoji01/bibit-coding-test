package domain

type Movie struct {
	Title     string `json:"Title"`
	Year      string `json:"Year"`
	ImdbID    string `json:"imdbID"`
	MovieType string `json:"Type"`
	Poster    string `json:"Poster"`

	Rated      string        `json:"Rated,omitempty"`
	Ratings    []MovieRating `json:"Ratings,omitempty"`
	Runtime    string        `json:"Runtime,omitempty"`
	Genre      string        `json:"Genre,omitempty"`
	Director   string        `json:"Director,omitempty"`
	Writer     string        `json:"Writer,omitempty"`
	Actors     string        `json:"Actors,omitempty"`
	Plot       string        `json:"Plot,omitempty"`
	Language   string        `json:"Language,omitempty"`
	Country    string        `json:"Country,omitempty"`
	Awards     string        `json:"Awards,omitempty"`
	Metascore  string        `json:"Metascore,omitempty"`
	ImdbRating string        `json:"imdbRating,omitempty"`
	ImdbVotes  string        `json:"imdbVotes,omitempty"`
	DVD        string        `json:"DVD,omitempty"`
	BoxOffice  string        `json:"BoxOffice,omitempty"`
	Production string        `json:"Production,omitempty"`
	Website    string        `json:"Website,omitempty"`
	Response   string        `json:"Response,omitempty"`
	Error      string        `json:"Error,omitempty"`
}

type MovieRating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type MovieUsecase interface {
	Fetch(query string, pagination int64) ([]Movie, error)
	GetByID(id string) (Movie, error)
}
