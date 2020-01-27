package main

import (
	"log"
	"net/http"

	poker "github.com/nhoffmann/learn-go-with-tests/command-line"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatal("could not listen on port 5000")
	}
}
