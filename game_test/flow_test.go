package game_test

import (
	"blackjack/blackjackErrors"
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/gameSessionState"
	"blackjack/hand"
	"blackjack/player"
	"github.com/adrianbrad/go-deck-of-cards"
	"testing"
)

func TestBetState(t *testing.T) {
	var g game.Game
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)

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
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)
	err := g.Bet(10)
	equals(t, g.GetState(), gameSessionState.StateBet)

	err = g.Bet(20)
	equals(t, err.Error(), blackjackErrors.BetAlreadyPlaced)

	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)

	err = g.Bet(31)
	equals(t, err.Error(), blackjackErrors.BetHigherThanBalance)

	err = g.Bet(-1)
	equals(t, err.Error(), blackjackErrors.BetHasToBeGreaterThanZero)
}

func TestDealAction(t *testing.T) {
	var g game.Game
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)
	_ = g.Bet(10)

	err := g.DealStartingHands()
	equals(t, err, nil)
}

func TestPlayerTurnState(t *testing.T) {
	var g game.Game
	invalidDeckForInsuranceAndSplitting := deck.Deck{{0, 5}, {0, 9}, {0, 7}, {0, 9}, {0, 3}}
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), invalidDeckForInsuranceAndSplitting)
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

	err = g.PlaceInsurance()
	equals(t, err.Error(), blackjackErrors.DealerFirstCardError)

	err = g.Split()
	//fmt.Println(err)
	equals(t, err.Error(), blackjackErrors.SplitCardsValueError)

	err = g.Hit()
	equals(t, err, nil)
	equals(t, len(g.GetPlayer().GetCurrentHandCards()), 3)

	err = g.Stand()
	equals(t, g.GetState(), gameSessionState.StateHandOver)

	validDeckForInsuranceNoBlackjackDealer := deck.Deck{{0, 5}, {0, 1}, {0, 3}, {0, 10}, {0, 3}}
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), validDeckForInsuranceNoBlackjackDealer)

	_ = g.Bet(10)

	_ = g.DealStartingHands()

	err = g.PlaceInsurance()
	equals(t, err, nil)
}

func TestEndStateSplit(t *testing.T) {
	var g game.Game
	validDeckForSplitting := deck.Deck{{1, 5}, {0, 9}, {2, 5}, {0, 4}, {0, 3}, {0, 6}}
	g = game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), validDeckForSplitting)

	err := g.Bet(10)
	equals(t, err, nil)

	err = g.DealStartingHands()
	equals(t, err, nil)

	err = g.Split()
	equals(t, err, nil)

	equals(t, g.GetPlayer().GetCurrentHandIndex(), uint8(1))
	equals(t, g.GetPlayer().GetCurrentHandBet(), 10)
	equals(t, g.GetPlayer().GetBalance(), 10)
	equals(t, g.GetPlayer().GetCurrentHandCards(), hand.Hand{{2, 5}})

	err = g.GetPlayer().SetCurrentIndexHand(uint8(0))
	equals(t, g.GetPlayer().GetCurrentHandIndex(), uint8(0))
	equals(t, g.GetPlayer().GetCurrentHandBet(), 10)
	equals(t, g.GetPlayer().GetCurrentHandCards(), hand.Hand{{1, 5}})
}
