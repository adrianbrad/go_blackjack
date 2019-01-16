package main

import (
	"blackjack/client"
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/player"
)

func main() {
	var g game.Game
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)

	var cliC client.Client
	cliC = client.NewCLIClient(g)

	cliC.DisplayOptions()
}
