package main

import (
	"blackjack/client"
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/player"
)

func main() {
	var g game.Game
	ga := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)
	
	g = &ga
	var cliC client.Client
	cliC = client.NewCLIClient(g)

	cliC.DisplayOptions()
}
