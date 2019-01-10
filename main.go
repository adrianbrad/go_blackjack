package main

import (
	"blackjack/game"
	"fmt"
)

func main() {
	var g game.Game
	g.Dealer = game.BasicDealer{}
	g.ShuffleNewDeck()
	fmt.Println(g.Deck)
	g.DealStartingHands()
	playerIndividualCardScore, playerTotalScore := g.PlayerHand.Score()
	dealerIndividualCardScore, dealerTotalScore := g.DealerHand.Score()
	fmt.Println("Starting Player Hand: ", g.PlayerHand, " Score: ", playerTotalScore, " Individual Score: ", playerIndividualCardScore)
	fmt.Println("Starting Dealer Hand: ", g.DealerHand, " Score: ", dealerTotalScore, " Individual Score: ", dealerIndividualCardScore)
	//humanMock := ai.HumanAI{}
	for g.State == game.StatePlayerTurn {
		//decision := humanMock.Play(g.PlayerHand, g.DealerHand[0])
		g.Hit()
		//decision(g)
		g.Stand()
	}

	if g.State == game.StateDealerTurn {
		g.FinishDealerHand()
		g.Stand()
	}

	playerIndividualCardScore, playerTotalScore = g.PlayerHand.Score()
	dealerIndividualCardScore, dealerTotalScore = g.DealerHand.Score()

	fmt.Println("Final Player Hand: ", g.PlayerHand, " Score: ", playerTotalScore, " Individual Score: ", playerIndividualCardScore)
	fmt.Println("Final Dealer Hand: ", g.DealerHand, " Score: ", dealerTotalScore, " Individual Score: ", dealerIndividualCardScore)
	fmt.Println(g.Deck)
	g.EndHand()
}
