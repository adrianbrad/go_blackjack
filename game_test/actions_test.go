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
	var g game.Game
	g = game.New(3, 1.5, player.New(50), dealer.NewDefaultDealer())

	g.ShuffleNewDeck()
	initialDeck := g.GetDeck()

	g.DealStartingHands()

	equals(t, g.GetState(), gameSessionState.StatePlayerTurn)
	g.Hit()
	g.Stand()
	equals(t, g.GetState(), gameSessionState.StateDealerTurn)
	equals(t, len(g.GetDeck()), 52*3-5)
	equals(t, g.GetDeck()[0], initialDeck[5])
	equals(t, g.GetPlayer().GetCurrentHandCards(), append(hand.Hand{initialDeck[0], initialDeck[2]}, initialDeck[4]))
}

func TestDealerHitThenStand(t *testing.T) {
	var g game.Game
	g = game.New(3, 1.5, player.New(50), dealer.NewDefaultDealer())

	g.ShuffleNewDeck()
	initialDeck := g.GetDeck()

	g.DealStartingHands()
	g.Hit()
	g.Stand() //set dealer turn

	equals(t, g.GetState(), gameSessionState.StateDealerTurn)

	g.Hit()
	g.Stand()

	equals(t, g.GetState(), gameSessionState.StateHandOver)
	equals(t, len(g.GetDeck()), 52*3-6)
	equals(t, g.GetDeck()[0], initialDeck[6])
	equals(t, g.GetDealer().GetDealerHand(), append(hand.Hand{initialDeck[1], initialDeck[3]}, initialDeck[5]))
}

func TestDoubleDown(t *testing.T) {
	var g game.Game
	g = game.New(3, 1.5, player.New(50), dealer.NewDefaultDealer())

	g.ShuffleNewDeck()
	g.Bet(10)
	equals(t, g.GetPlayer().GetCurrentHandBet(), 10)

	g.DealStartingHands()
	equals(t, g.GetState(), gameSessionState.StatePlayerTurn)

	g.DoubleDown()

	equals(t, g.GetPlayer().GetCurrentHandBet(), 20)
	equals(t, g.GetState(), gameSessionState.StateDealerTurn)
	equals(t, len(g.GetPlayer().GetCurrentHandCards()), 3)
}
