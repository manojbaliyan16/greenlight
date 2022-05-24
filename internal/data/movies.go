package data

import "time"

//1. first time taken this simple struct without JSON tag
/*type Movie struct {
	ID        int64
	CreatedAt time.Time // timestamp when movie is added
	Title     string
	Year      int32
	Runtime   int32
	Genres    []string // slice of genres of movies // (romance, comedy, etc)
	Version   int32
}*/

// 2. Annotate the movies with struct tags to control how the keys appear in the JSON -encoded output

/*type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"` // timestamp when movie is added
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Runtime   int32     `json:"runtime"`
	Genres    []string  `json:"genres"` // slice of genres of movies // (romance, comedy, etc)
	Version   int32     `json:"version"`
}*/
// 3. Hides the particular field in response

// Movie endpoint example http://localhost:4000/v1/movies/123
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // timestamp when movie is added
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"` // slice of genres of movies // (romance, comedy, etc)
	Version   int32     `json:"version"`
}
