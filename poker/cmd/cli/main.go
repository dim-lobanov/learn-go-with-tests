package main

import (
	"fmt"
	"log"
	"os"
	"learn-go-with-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
