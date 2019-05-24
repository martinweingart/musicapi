//Package models defines the model of the API
package models

//Artist data model
type Artist struct {
	Name   string  `json:"name"`
	Albums []Album `json:"albums"`
}

//ArtistSimple is the model with reduced data of the artist (withouth Albums)
type ArtistSimple struct {
	Name string `json:"name"`
}

//Album data model
type Album struct {
	Name   string  `json:"name"`
	Artist string  `json:"artist"`
	Year   string  `json:"year"`
	Tracks []Track `json:"tracks"`
}

//AlbumSimple is the model with reduced data of the album (withouth Tracks)
type AlbumSimple struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Year   string `json:"year"`
}

//Track data model
type Track struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	File   string `json:"file"`
}
