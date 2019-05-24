//Package dblayer is used to perform database operations
//This file implements the database layer on sqlite
package dblayer

import (
	"database/sql"
	"log"
	"musicapi/models"
	str "strings"

	//SQLite3 driver package
	_ "github.com/mattn/go-sqlite3"
)

//DB holds a reference of the db connection
type DB struct {
	*sql.DB
}

//MP3Data is a class of object with the mp3 tag info
type MP3Data struct {
	Artist string
	Album  string
	Year   string
	Genre  string
	Title  string
	File   string
	Update int64
}

//NewDB creates a new database connection. Will create the database file if not exists
func NewDB(dbname, con string) (*DB, error) {
	db, err := sql.Open(dbname, con)
	if err != nil {
		log.Fatal("Error initializing databse!", err)
		return nil, nil
	}
	return &DB{
		DB: db,
	}, err
}

//Init will initialize the database schema
func (db *DB) Init() {
	log.Println("Creating database schema...")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS mp3_file (file VARCHAR(255) PRIMARY KEY, artist VARCHAR(100), album VARCHAR(100), year VARCHAR(4), genre VARCHAR(100), title VARCHAR(100), lastUpdate VARCHAR(20))")
	statement.Exec()
	statement.Close()
}

//SaveMP3Data inserts the mp3 info into the database
func (db *DB) SaveMP3Data(data MP3Data) {
	statement, err := db.Prepare("INSERT INTO mp3_file (file, artist, album, year, genre, title, lastUpdate) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Save mp3 info: error preparing insert query")
		return
	}

	_, err = statement.Exec(data.File, data.Artist, data.Album, data.Year, data.Genre, data.Title, data.Update)

	if err != nil {
		if str.Contains(err.Error(), "UNIQUE constraint failed") {
			statement, err := db.Prepare("UPDATE mp3_file SET lastUpdate=? WHERE file=?")
			if err != nil {
				log.Println("Save mp3 info: error preparing update query")
				return
			}

			_, err = statement.Exec(data.Update, data.File)
			if err != nil {
				log.Println("Error updating mp3 info")
			}
		} else {
			log.Println("Error inserting mp3 info")
		}

		return
	}

	statement.Close()
}

//Clean removes the rows that contains info of files that doesn't exists anymore
func (db *DB) Clean(lastUpdate int64) {
	log.Println("Removing deleted files from the database...")
	statement, _ := db.Prepare("DELETE FROM mp3_file WHERE lastUpdate != ?")
	statement.Exec(lastUpdate)
	statement.Close()
}

//GetAllArtists Get all artists
func (db *DB) GetAllArtists() (artists []models.ArtistSimple, err error) {
	rows, err := db.Query("SELECT DISTINCT artist FROM mp3_file")
	if err != nil {
		log.Println("Error executing SQL query")
		return nil, err
	}

	for rows.Next() {
		var name string
		rows.Scan(&name)
		artist := models.ArtistSimple{
			Name: name,
		}
		artists = append(artists, artist)
	}
	return artists, err
}

//GetArtist Get an artist based on his name
func (db *DB) GetArtist(name string) (artist models.Artist, err error) {
	statement, err := db.Prepare("SELECT DISTINCT artist, album, year FROM mp3_file WHERE artist=?")
	rows, err := statement.Query(name)

	var artistName, album, year string
	artist = models.Artist{
		Name:   name,
		Albums: []models.Album{},
	}

	for rows.Next() {
		rows.Scan(&artistName, &album, &year)
		idAlbum := name + album + year
		log.Println(idAlbum)
		album, _ := db.GetAlbum(idAlbum)
		artist.Albums = append(artist.Albums, album)
	}

	return artist, err
}

//GetAllAlbums Get all albums
func (db *DB) GetAllAlbums() (albums []models.AlbumSimple, err error) {
	rows, err := db.Query("SELECT DISTINCT album, artist, year FROM mp3_file")
	for rows.Next() {
		var name, artist, year string
		rows.Scan(&name, &artist, &year)
		album := models.AlbumSimple{
			Name:   name,
			Artist: artist,
			Year:   year,
		}
		albums = append(albums, album)
	}
	return albums, err
}

//GetAlbum Get an album based on a id formed as the concatination of the artist, album name and year
func (db *DB) GetAlbum(id string) (album models.Album, err error) {
	statement, err := db.Prepare("SELECT artist, album, title, genre, year, file FROM mp3_file WHERE artist || album || year=?")

	rows, err := statement.Query(id)

	var artist, albumName, title, genre, year, file string
	rows.Next()
	rows.Scan(&artist, &albumName, &title, &genre, &year, &file)

	album = models.Album{
		Name:   albumName,
		Artist: artist,
		Year:   year,
		Tracks: []models.Track{},
	}

	track := models.Track{
		Title:  title,
		Artist: artist,
		Album:  albumName,
		Genre:  genre,
		Year:   year,
		File:   file,
	}
	album.Tracks = append(album.Tracks, track)

	for rows.Next() {
		rows.Scan(&artist, &albumName, &title, &genre, &year, &file)
		track = models.Track{
			Title:  title,
			Artist: artist,
			Album:  albumName,
			Genre:  genre,
			Year:   year,
			File:   file,
		}
		album.Tracks = append(album.Tracks, track)
	}

	return album, err
}

//GetAllTracks Get all the mp3 files (tracks)
func (db *DB) GetAllTracks() (tracks []models.Track, err error) {
	rows, err := db.Query("SELECT title, artist, album, genre, year, file FROM mp3_file")
	for rows.Next() {
		var title, artist, album, genre, year, file string
		rows.Scan(&title, &artist, &album, &genre, &year, &file)
		track := models.Track{
			Title:  title,
			Artist: artist,
			Album:  album,
			Genre:  genre,
			Year:   year,
			File:   file,
		}
		tracks = append(tracks, track)
	}
	return tracks, err
}
