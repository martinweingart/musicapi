//Package dblayer is used to perform database operations
package dblayer

import "musicapi/models"

//DBLayer interface that defines the methods for retrieving information of the database
type DBLayer interface {
	GetAllArtists() ([]models.ArtistSimple, error)
	GetArtist(string) (models.Artist, error)
	GetAllAlbums() ([]models.AlbumSimple, error)
	GetAlbum(string) (models.Album, error)
	GetAllTracks() ([]models.Track, error)
}
