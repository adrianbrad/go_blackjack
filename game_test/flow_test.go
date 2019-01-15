package game_test

import (
	"blackjack/blackjackErrors"
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/gameSessionState"
	"blackjack/player"
	"testing"
)

func TestBetState(t *testing.T) {
	var g game.Game
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer())

	err := g.Hit()
	equals(t, err.Error(), blackjackErrors.NoActiveHands)

	err = g.Stand()
	equals(t, err.Error(), blackjackErrors.NoActiveHands)

	err = g.DoubleDown()
	equals(t, err.Error(), "given state: StatePlayerTurn different from the current state: StateBet")

	err = g.Split()
	equals(t, err.Error(), "given state: StatePlayerTurn different from the current state: StateBet")

	err = g.PlaceInsurance()
	equals(t, err.Error(), "given state: StatePlayerTurn different from the current state: StateBet")

	err = g.FinishDealerHand()
	equals(t, err.Error(), "given state: StateDealerTurn different from the current state: StateBet")

	err = g.DealStartingHands()
	equals(t, err.Error(), blackjackErrors.NoBetsPlaced)

	_, _, _, err = g.EndHand()
	equals(t, err.Error(), "given state: StateHandOver different from the current state: StateBet")

	equals(t, g.GetState(), gameSessionState.StateBet)
}

func TestBetAction(t *testing.T) {
	var g game.Game
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer())
	err := g.Bet(10)
	equals(t, g.GetState(), gameSessionState.StateBet)

	err = g.Bet(20)
	equals(t, err.Error(), blackjackErrors.BetAlreadyPlaced)

	err = g.Bet(31)
	equals(t, err.Error(), blackjackErrors.BetHigherThanBalance)

	err = g.Bet(-1)
	equals(t, err.Error(), blackjackErrors.BetHasToBeGreaterThanZero)
}

func TestDealAction(t *testing.T) {
	var g game.Game
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer())
	_ = g.Bet(10)

	err := g.DealStartingHands()
	equals(t, err, nil)
}

func TestPlayerTurnState(t *testing.T) {
	var g game.Game
	//invalidDeckForInsuranceAndSplitting := deck.Deck{{0,5},{1,6}, {2,7}} //TODO finish this
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer())
	_ = g.Bet(10)

	_ = g.DealStartingHands()

	err := g.Bet(10)
	equals(t, err.Error(), "given state: StateBet different from the current state: StatePlayerTurn")

	err = g.DealStartingHands()
	equals(t, err.Error(), "given state: StateBet different from the current state: StatePlayerTurn")

	err = g.FinishDealerHand()
	equals(t, err.Error(), "given state: StateDealerTurn different from the current state: StatePlayerTurn")

	_, _, _, err = g.EndHand()
	equals(t, err.Error(), "given state: StateHandOver different from the current state: StatePlayerTurn")

	//err = g.PlaceInsurance()
	//equals(t, err.Error(), "given state: StateHandOver different from the current state: StatePlayerTurn")

}
