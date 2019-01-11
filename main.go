package main

import (
	"blackjack/game"
	"fmt"
)

func main() {
	g := game.Game{
		Dealer:          game.BasicDealer{},
		State:           game.StatePlayerTurn,
		NumDecks:        3,
		BlackjackPayout: 1.5,
		PlayerBalance:   50,
	}

	g.ShuffleNewDeck()

	g.DealStartingHands()

	playerIndividualCardScore := g.PlayerHand.IndividualScore()
	playerTotalScore := g.PlayerHand.Score()
	dealerIndividualCardScore := g.DealerHand.IndividualScore()
	dealerTotalScore := g.DealerHand.Score()
	fmt.Println("Starting Player Hand: ", g.PlayerHand, " Score: ", playerTotalScore, " Individual Score: ", playerIndividualCardScore)
	fmt.Println("Starting Dealer Hand: ", g.DealerHand, " Score: ", dealerTotalScore, " Individual Score: ", dealerIndividualCardScore)
	//humanMock := ai.HumanAI{}
	g.Bet(10)
	for g.State == game.StatePlayerTurn {
		//decision := humanMock.Play(g.PlayerHand, g.DealerHand[0])
		//g.DoubleDown()
		g.Hit()
		//decision(g)
		g.Stand()
	}

	if g.State == game.StateDealerTurn {
		g.FinishDealerHand()
		g.Stand()
	}

	playerIndividualCardScore = g.PlayerHand.IndividualScore()
	playerTotalScore = g.PlayerHand.Score()
	dealerIndividualCardScore = g.DealerHand.IndividualScore()
	dealerTotalScore = g.DealerHand.Score()

	fmt.Println("Final Player Hand: ", g.PlayerHand, " Score: ", playerTotalScore, " Individual Score: ", playerIndividualCardScore)
	fmt.Println("Final Dealer Hand: ", g.DealerHand, " Score: ", dealerTotalScore, " Individual Score: ", dealerIndividualCardScore)
	fmt.Println(g.Deck)
	g.EndHand()
}
