package player

import (
	"blackjack/blackjackErrors"
	"blackjack/hand"
	"fmt"
	"github.com/adrianbrad/go-deck-of-cards"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	var p Player
	p = New(10)
	equals(t, p.GetBalance(), 10)
	equals(t, p.GetCurrentHandIndex(), uint8(0))
	equals(t, p.GetTotalHands(), uint8(1))
	equals(t, len(p.GetCurrentHandCards()), 0)
	equals(t, p.GetCurrentHandBet(), 0)
}

func TestPlayer_SetCurrentHandBet(t *testing.T) {
	var p Player
	p = New(20)
	err := p.SetCurrentHandBet(5)
	equals(t, p.GetCurrentHandBet(), 5)
	equals(t, p.GetBalance(), 15)
	equals(t, err, nil)

	err = p.SetCurrentHandBet(17)
	equals(t, err, fmt.Errorf(blackjackErrors.BetAlreadyPlaced))

	p.ResetHands()
	p.SetBalance(50)

	err = p.SetCurrentHandBet(0)
	equals(t, err, fmt.Errorf(blackjackErrors.BetHasToBeGreaterThanZero))
	equals(t, p.GetCurrentHandBet(), 0)
	equals(t, p.GetBalance(), 50)

	err = p.SetCurrentHandBet(51)
	equals(t, err, fmt.Errorf(blackjackErrors.BetHigherThanBalance))
	equals(t, p.GetCurrentHandBet(), 50)
	equals(t, p.GetBalance(), 0)

}

func TestPlayer_DoubleCurrentHandBet(t *testing.T) {
	var p Player
	p = New(20)

	_ = p.SetCurrentHandBet(5)
	p.GetCurrentHandCardsPointer().AddCard(deck.Card{1, 9})
	p.GetCurrentHandCardsPointer().AddCard(deck.Card{1, 9})

	err := p.DoubleCurrentHandBet()
	equals(t, err, nil)
	equals(t, p.GetCurrentHandBet(), 10)
	equals(t, p.GetBalance(), 10)

	err = p.DoubleCurrentHandBet()
	equals(t, err, fmt.Errorf("bet already doubled"))

	p.ResetHands()
	err = p.DoubleCurrentHandBet()
	equals(t, err, fmt.Errorf(blackjackErrors.NoBetsPlaced))

	p.SetBalance(20)
	_ = p.SetCurrentHandBet(15)
	err = p.DoubleCurrentHandBet()
	equals(t, err, fmt.Errorf(blackjackErrors.BetHigherThanBalance))
	equals(t, p.GetCurrentHandBet(), 20)
	equals(t, p.GetBalance(), 0)
}

func TestPlayer_SplitHands(t *testing.T) {
	var p Player
	p = New(50)

	_ = p.SetCurrentHandBet(5)
	(*p.GetCurrentHandCardsPointer()).AddCard(deck.Card{deck.Spades, deck.Ace})
	(*p.GetCurrentHandCardsPointer()).AddCard(deck.Card{deck.Clubs, deck.Ace})

	err := p.SplitHands()
	equals(t, err, nil)
	equals(t, p.GetTotalHands(), uint8(2))

	equals(t, p.GetBalance(), 40)

	equals(t, p.GetCurrentHandIndex(), uint8(1))
	equals(t, p.GetCurrentHandBet(), 5)
	equals(t, p.GetCurrentHandCards(), hand.Hand{{deck.Clubs, deck.Ace}})
	equals(t, len(p.GetCurrentHandCards()), 1)

	err = p.SetCurrentIndexHand(0)
	equals(t, err, nil)

	equals(t, p.GetTotalHands(), uint8(2))

	equals(t, p.GetBalance(), 40)

	equals(t, p.GetCurrentHandIndex(), uint8(0))
	equals(t, p.GetCurrentHandBet(), 5)
	equals(t, p.GetCurrentHandCards(), hand.Hand{{deck.Spades, deck.Ace}})
	equals(t, len(p.GetCurrentHandCards()), 1)

}

func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
