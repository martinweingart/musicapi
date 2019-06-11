package main

import (
	"log"
	"flag"
	"musicapi/dblayer"
	"musicapi/rest"
	"musicapi/scan"
	"os"
)

func main() {
	folder := flag.String("path", "", "Specify the path of the music folder to scan")

	log.Println("Looking for an existing database file...")

	db, err := dblayer.NewDB("sqlite3", "music.db")
	if err != nil {
		log.Fatal("Error initializing database")
		return
	}

	if _, err := os.Stat("music.db"); os.IsNotExist(err) {
		db.Init()
	}

	if *folder != "" {
		scan.Scan(db, *folder)
	}

	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
