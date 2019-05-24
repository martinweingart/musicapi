package main

import (
	"log"
	"musicapi/dblayer"
	"musicapi/rest"
	"musicapi/scan"
	"os"
)

func main() {
	folder := os.Args[1]
	log.Println("Looking for an existing database file...")

	db, err := dblayer.NewDB("sqlite3", "music.db")
	if err != nil {
		log.Fatal("Error initializing database")
		return
	}

	if _, err := os.Stat("music.db"); os.IsNotExist(err) {
		db.Init()
	}

	scan.Scan(db, folder)
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
