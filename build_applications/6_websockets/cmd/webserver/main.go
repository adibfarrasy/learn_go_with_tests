package main

import (
	"log"
	"net/http"
	"server/poker"
	. "server/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

    game := poker.NewTexasHoldemGame(poker.BlindAlerterFunc(poker.Alerter), store)
	server, _ := NewPlayerServer(store, game)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen to port 5000 %v", err)
	}
}
