package game_test

import (
	"blackjack/game"
	"blackjack/hand"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestDealStartingHands(t *testing.T) {
	g := game.Game{
		Dealer:          game.BasicDealer{},
		State:           game.StatePlayerTurn,
		NumDecks:        3,
		BlackjackPayout: 1.5,
		PlayerBalance:   50,
	}

	g.ShuffleNewDeck()
	initialDeck := g.Deck

	equals(t, len(g.Deck), 52*3)

	g.DealStartingHands()

	equals(t, len(g.PlayerHand), 2)
	equals(t, len(g.DealerHand), 2)

	equals(t, g.PlayerHand, append(hand.Hand{initialDeck[0]}, initialDeck[2]))
	equals(t, g.DealerHand, append(hand.Hand{initialDeck[1]}, initialDeck[3]))

	equals(t, len(g.Deck), 52*3-4)
	equals(t, g.Deck[0], initialDeck[4])
}

func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
