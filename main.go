package main

import (
	"blackjack/game"
	"blackjack/player"
	"fmt"
)

func main() {

	g := game.Game{
		Dealer:          game.BasicDealer{},
		State:           game.StatePlayerTurn,
		NumDecks:        3,
		BlackjackPayout: 1.5,
		Player:          player.New(50),
	}

	g.ShuffleNewDeck()
	fmt.Println(g.Deck)
	g.DealStartingHands()

	playerIndividualCardScore := g.Player.GetCurrentHandCards().IndividualScore()
	playerTotalScore := g.Player.GetCurrentHandCards().Score()
	dealerIndividualCardScore := g.DealerHand.IndividualScore()
	dealerTotalScore := g.DealerHand.Score()
	fmt.Println("Starting Player Hand: ", g.Player.GetCurrentHandCards(), " Score: ", playerTotalScore, " Individual Score: ", playerIndividualCardScore)
	fmt.Println("Starting Dealer Hand: ", g.DealerHand, " Score: ", dealerTotalScore, " Individual Score: ", dealerIndividualCardScore)
	//humanMock := ai.HumanAI{}
	g.Bet(10)
	g.DoubleDown()
	//for g.State == game.StatePlayerTurn {
	//	//decision := humanMock.Play(g.PlayerHand, g.DealerHand[0])
	//	//g.DoubleDown()
	//	g.Hit()
	//	//decision(g)
	//	g.Stand()
	//}

	if g.State == game.StateDealerTurn {
		g.FinishDealerHand()
		g.Stand()
	}

	playerIndividualCardScore = g.Player.GetCurrentHandCards().IndividualScore()
	playerTotalScore = g.Player.GetCurrentHandCards().Score()
	dealerIndividualCardScore = g.DealerHand.IndividualScore()
	dealerTotalScore = g.DealerHand.Score()

	fmt.Println("Final Player Hand: ", g.Player.GetCurrentHandCards(), " Score: ", playerTotalScore, " Individual Score: ", playerIndividualCardScore)
	fmt.Println("Final Dealer Hand: ", g.DealerHand, " Score: ", dealerTotalScore, " Individual Score: ", dealerIndividualCardScore)
	fmt.Println(g.Deck)
	g.EndHand()
}
