package main

import (
	"blackjack/game"
	"fmt"
)

func main() {
	var g game.Game
	g.ShuffleNewDeck()
	g.DealStartingHands()
	fmt.Println("Player Hand: ", g.PlayerHand, " Score: ", g.PlayerHand.Score())
	fmt.Println("Dealer Hand: ", g.DealerHand, " Score: ", g.DealerHand.Score())
	//switch "a" {
	//case "hit":
	g.Hit()
	//case "stand":
	g.Stand()
	//}
	if g.State == game.StateDealerTurn {
		g.FinishDealerHand()
		g.Stand()
	}

	fmt.Println("Player Hand: ", g.PlayerHand, " Score: ", g.PlayerHand.Score())
	fmt.Println("Dealer Hand: ", g.DealerHand, " Score: ", g.DealerHand.Score())

	g.EndHand()
}
