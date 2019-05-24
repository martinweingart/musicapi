//Package scan is used to scan a folder searching mp3 files and saving tag data
package scan

import (
	"log"
	"musicapi/dblayer"
	"os"
	"path/filepath"
	"strings"
	"time"

	id3 "github.com/mikkyang/id3-go"
)

//Creates an object MP3Data with the mp3 tag data and calls the database layer for saving it
func saveTags(db *dblayer.DB, path string, mp3 *id3.File, stamp int64) {
	data := dblayer.MP3Data{
		Artist: mp3.Artist(),
		Album:  mp3.Album(),
		Year:   mp3.Year(),
		Title:  mp3.Title(),
		File:   path,
		Update: stamp,
	}
	db.SaveMP3Data(data)
}

//Check if a file is an mp3 one
func checkMP3Ext(path string) bool {
	return filepath.Ext(strings.TrimSpace(path)) == ".mp3"
}

//Scan will scan a folder searching mp3 files and saving their information
func Scan(db *dblayer.DB, folder string) {
	log.Println("Starting scanning of folder: ", folder)
	stamp := time.Now().Unix()
	err := filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		mp3File, err := id3.Open(path)
		if err == nil && checkMP3Ext(path) {
			saveTags(db, path, mp3File, stamp)
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error scanning folder")
		return
	}

	db.Clean(stamp)
	log.Println("Folder scanning finished!")
}
