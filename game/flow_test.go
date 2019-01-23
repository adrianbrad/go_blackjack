package game_test

import (
	"blackjack/blackjackErrors"
	"blackjack/blackjackWinner"
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/gameSessionState"
	"blackjack/hand"
	"blackjack/outcome"
	"blackjack/player"
	"testing"

	deck "github.com/adrianbrad/go-deck-of-cards"
)

func TestBetState(t *testing.T) {
	g := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)

	err := g.Hit()
	equals(t, err.Error(), blackjackErrors.HitPlayerTurnError)

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
	equals(t, err.Error(), "given state: StateHandOver different from the current state: StatePlayerTurn")

	equals(t, g.GetState(), gameSessionState.StatePlayerTurn)
}

func TestBetAction(t *testing.T) {
	g := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)

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
	g := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), nil)

	g.Bet(10)
	err := g.DealStartingHands()
	equals(t, err, nil)
}

func TestPlayerTurnState(t *testing.T) {
	invalidDeckForInsuranceAndSplitting := deck.Deck{{0, 5}, {0, 9}, {0, 7}, {0, 9}, {0, 3}, {1, 5}, {2, 5}}
	g := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), invalidDeckForInsuranceAndSplitting)

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

func TestEndState(t *testing.T) {
	playerBlackjackWinDeck := deck.Deck{{1, 1}, {0, 9}, {2, 10}, {0, 9}}
	g := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), playerBlackjackWinDeck)

	_ = g.Bet(10)

	_ = g.DealStartingHands()

	_ = g.Stand()

	outcomes, playerWinnings, moneyOperations, err := g.EndHand()
	equals(t, err, nil)

	equals(t, outcomes, []outcome.BlackjackOutcome{outcome.NewBlackjackOutcome(0, false, true)})

	equals(t, playerWinnings, outcome.Winnings(25))

	mo, _ := outcome.ComputeMoneyOperationsForHand(10, outcome.BetBack(10), outcome.Winnings(15))
	moneyOperationsExpected := [][]outcome.MoneyOperation{mo}
	
	equals(t, moneyOperations, moneyOperationsExpected)
}

func TestEndStateSplitPlayerWinsBothHands(t *testing.T) {
	validDeckForSplitting := deck.Deck{{1, 5}, {0, 9}, {2, 5}, {0, 9}, {0, 11}, {0, 4}, {0, 5}, {0, 1}}
	g := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), validDeckForSplitting)

	err := g.Bet(10)
	equals(t, err, nil)

	err = g.DealStartingHands()
	equals(t, err, nil)

	err = g.Split()
	equals(t, err, nil)

	equals(t, g.GetPlayer().GetCurrentHandIndex(), uint8(1))
	equals(t, g.GetPlayer().GetCurrentHandBet(), 10)
	equals(t, g.GetPlayer().GetBalance(), 10)
	equals(t, g.GetPlayer().GetCurrentHandCards(), hand.Hand{{2, 5}, {0, 11}})

	err = g.GetPlayer().SetCurrentHandIndex(uint8(0))
	equals(t, g.GetPlayer().GetCurrentHandIndex(), uint8(0))
	equals(t, g.GetPlayer().GetCurrentHandBet(), 10)
	equals(t, g.GetPlayer().GetCurrentHandCards(), hand.Hand{{1, 5}})

	g.GetPlayer().SetCurrentHandIndex(uint8(1))

	g.Hit()

	equals(t, g.GetPlayer().GetCurrentHandCards(), hand.Hand{{2, 5}, {0, 11}, {0, 4}})

	g.Stand()

	equals(t, g.GetPlayer().GetCurrentHandIndex(), uint8(0))

	equals(t, g.GetPlayer().GetCurrentHandCards(), hand.Hand{{1, 5}, {0, 5}})

	g.Hit()

	equals(t, g.GetPlayer().GetCurrentHandCards(), hand.Hand{{1, 5}, {0, 5}, {0, 1}})

	g.Stand()

	equals(t, g.GetState(), gameSessionState.StateHandOver)

	outcomes, winnings, _, err := g.EndHand()

	equals(t, err, nil)

	equals(t, outcomes[0], outcome.NewBlackjackOutcome(blackjackWinner.Player, false, false))

	equals(t, outcomes[1], outcome.NewBlackjackOutcome(blackjackWinner.Player, false, false))

	equals(t, winnings, outcome.Winnings(20+20))
	equals(t, g.GetPlayer().GetBalance(), 30-10-10+20+20)
	// equals(t, operations, nil)
}

func TestEndStateSplitPlayerWinsOneHandWithBlackjackLosesTheOtherWithBust(t *testing.T) {
	validDeckForSplitting := deck.Deck{{1, 10}, {0, 9}, {2, 10}, {0, 9}, {0, 1}, {0, 6}, {0, 10}, {0, 1}}

	g := game.New(3, 1.5, player.New(30), dealer.NewDefaultDealer(), validDeckForSplitting)

	err := g.Bet(10)
	equals(t, err, nil)

	err = g.DealStartingHands()
	equals(t, err, nil)

	err = g.Split()
	equals(t, err, nil)

	// on the second hand we have a blackjack so we jump back to the first hand
	equals(t, g.GetPlayer().GetCurrentHandIndex(), uint8(0))

	equals(t, len(g.GetPlayer().GetCurrentHandCards()), 2)

	g.Hit()

	equals(t, len(g.GetPlayer().GetCurrentHandCards()), 3)

	err = g.Hit()

	equals(t, err.Error(), blackjackErrors.HitPlayerTurnError)

	outcomes, winnings, _, err := g.EndHand()

	equals(t, err, nil)
	equals(t, outcomes[1], outcome.NewBlackjackOutcome(blackjackWinner.Player, false, true))

	g.GetPlayer().SetCurrentHandIndex(0)
	equals(t, outcomes[0], outcome.NewBlackjackOutcome(blackjackWinner.Dealer, true, false))

	equals(t, winnings, outcome.Winnings(25+0))
	equals(t, g.GetPlayer().GetBalance(), 30-10-10+25)
}
