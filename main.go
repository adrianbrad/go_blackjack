package main

import (
	"blackjack/client"
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/player"
	"flag"
	"strconv"

	deck "github.com/adrianbrad/go-deck-of-cards"
)

func main() {
	predefinedDeckStringPtr := flag.String("deck", "", "the cards given")
	flag.Parse()

	var predefinedDeck deck.Deck

	for _, rune := range *predefinedDeckStringPtr {
		rank, _ := strconv.ParseInt(string(rune), 16, 8)
		predefinedDeck = append(predefinedDeck, deck.Card{0, deck.Rank(rank)})
	}

	g := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), predefinedDeck) //nil if the flag is not used and we generate a random one

	var cliC client.Client
	cliC = client.NewCLIClient(g)

	cliC.DisplayOptions()
}
