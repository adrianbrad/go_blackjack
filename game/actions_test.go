package game_test

import (
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/gameSessionState"
	"blackjack/hand"
	"blackjack/player"
	"testing"
)

func TestPlayerHitThenStand(t *testing.T) {
	g := game.New(3, 1.5, player.New(50), dealer.NewDefaultDealer(), nil)

	g.ShuffleNewDeck()
	initialDeck := g.GetDeck()
	g.Bet(10)
	_ = g.DealStartingHands()

	equals(t, g.GetState(), gameSessionState.StatePlayerTurn)
	g.Hit()
	g.Stand()
	equals(t, g.GetState(), gameSessionState.StateHandOver)
	equals(t, len(g.GetDeck()), 52*3-len(g.GetPlayer().GetCurrentHandCards())-len(g.GetDealer().GetDealerHand()))
	equals(t, g.GetDeck()[0], initialDeck[len(g.GetPlayer().GetCurrentHandCards())+len(g.GetDealer().GetDealerHand())])
	equals(t, g.GetPlayer().GetCurrentHandCards(), append(hand.Hand{initialDeck[0], initialDeck[2]}, initialDeck[4]))
}

func TestDoubleDown(t *testing.T) {
	g := game.New(3, 1.5, player.New(50), dealer.NewDefaultDealer(), nil)
	
	g.ShuffleNewDeck()
	g.Bet(10)
	equals(t, g.GetPlayer().GetCurrentHandBet(), 10)

	g.DealStartingHands()
	equals(t, g.GetState(), gameSessionState.StatePlayerTurn)

	g.DoubleDown()

	equals(t, g.GetPlayer().GetCurrentHandBet(), 20)
	equals(t, g.GetState(), gameSessionState.StateHandOver)
	equals(t, len(g.GetPlayer().GetCurrentHandCards()), 3)
}
