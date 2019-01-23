package game_test

import (
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/hand"
	"blackjack/player"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestDealStartingHands(t *testing.T) {
	g := game.New(3, 1.5, player.New(50), dealer.NewDefaultDealer(), nil)

	g.ShuffleNewDeck()
	initialDeck := g.GetDeck()

	equals(t, len(g.GetDeck()), 52*3)
	_ = g.Bet(10)
	_ = g.DealStartingHands()

	equals(t, len(g.GetPlayer().GetCurrentHandCards()), 2)
	equals(t, len(g.GetDealer().GetDealerHand()), 2)

	equals(t, g.GetPlayer().GetCurrentHandCards(), append(hand.Hand{initialDeck[0]}, initialDeck[2]))
	equals(t, g.GetDealer().GetDealerHand(), append(hand.Hand{initialDeck[1]}, initialDeck[3]))

	equals(t, len(g.GetDeck()), 52*3-4)
	equals(t, g.GetDeck()[0], initialDeck[4])
}

func equals(tb testing.TB, act, exp interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
