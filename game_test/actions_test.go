package game_test

import (
	"blackjack/game"
	"blackjack/hand"
	"testing"
)

func TestPlayerHitThenStand(t *testing.T) {
	g := game.Game{
		Dealer:          game.BasicDealer{},
		State:           game.StatePlayerTurn,
		NumDecks:        3,
		BlackjackPayout: 1.5,
		PlayerBalance:   50,
	}

	g.ShuffleNewDeck()
	initialDeck := g.Deck

	g.DealStartingHands()

	equals(t, g.State, game.StatePlayerTurn)
	g.Hit()
	g.Stand()
	equals(t, g.State, game.StateDealerTurn)
	equals(t, len(g.Deck), 52*3-5)
	equals(t, g.Deck[0], initialDeck[5])
	equals(t, g.PlayerHand, append(hand.Hand{initialDeck[0], initialDeck[2]}, initialDeck[4]))
}

func TestDealerHitThenStand(t *testing.T) {
	g := game.Game{
		Dealer:          game.BasicDealer{},
		State:           game.StatePlayerTurn,
		NumDecks:        3,
		BlackjackPayout: 1.5,
		PlayerBalance:   50,
	}

	g.ShuffleNewDeck()
	initialDeck := g.Deck

	g.DealStartingHands()
	g.Hit()
	g.Stand() //set dealer turn

	equals(t, g.State, game.StateDealerTurn)

	g.Hit()
	g.Stand()

	equals(t, g.State, game.StateHandOver)
	equals(t, len(g.Deck), 52*3-6)
	equals(t, g.Deck[0], initialDeck[6])
	equals(t, g.DealerHand, append(hand.Hand{initialDeck[1], initialDeck[3]}, initialDeck[5]))
}
