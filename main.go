package main

import (
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/player"
	"fmt"
)

func main() {
	var g game.Game
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer())

	g.ShuffleNewDeck()
	fmt.Println(g.GetDeck())
	g.DealStartingHands()

	playerIndividualCardScore := g.GetPlayer().GetCurrentHandCards().IndividualScore()
	playerTotalScore := g.GetPlayer().GetCurrentHandCards().Score()
	dealerIndividualCardScore := g.GetDealer().GetDealerHand().IndividualScore()
	dealerTotalScore := g.GetDealer().GetDealerHand().Score()
	fmt.Println("Starting Player Hand: ", g.GetPlayer().GetCurrentHandCards(), " Score: ", playerTotalScore, " Individual Score: ", playerIndividualCardScore)
	fmt.Println("Starting Dealer Hand: ", g.GetDealer().GetDealerFirstCard(), " Score: ", dealerTotalScore, " Individual Score: ", dealerIndividualCardScore)

	g.Bet(10)

	g.Hit()
	g.Stand()

	g.FinishDealerHand()

	playerIndividualCardScore = g.GetPlayer().GetCurrentHandCards().IndividualScore()
	playerTotalScore = g.GetPlayer().GetCurrentHandCards().Score()
	dealerIndividualCardScore = g.GetDealer().GetDealerHand().IndividualScore()
	dealerTotalScore = g.GetDealer().GetDealerHand().Score()

	fmt.Println("Final Player Hand: ", g.GetPlayer().GetCurrentHandCards(), " Score: ", playerTotalScore, " Individual Score: ", playerIndividualCardScore)
	fmt.Println("Final Dealer Hand: ", g.GetDealer().GetDealerHand(), " Score: ", dealerTotalScore, " Individual Score: ", dealerIndividualCardScore)
	fmt.Println(g.GetDeck())
	g.EndHand()
}
