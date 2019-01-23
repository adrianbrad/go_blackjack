package player

import (
	"blackjack/blackjackErrors"
	"blackjack/hand"
	"blackjack/outcome"
	"fmt"
)

type playerHand struct {
	cards      hand.Hand
	bet        int
	insurance  int
	doubleDown bool
	winnings   outcome.Winnings
	outcome    outcome.BlackjackOutcome
}

type player struct {
	currentHandIndex uint8
	totalHands       uint8
	balance          int
	hands            []playerHand
}

type Player interface {
	SetCurrentHandBet(int) error
	GetCurrentHandBet() int
	DoubleCurrentHandBet() error

	SetCurrentHandIndex(uint8) error
	GetCurrentHandIndex() uint8

	GetTotalHands() uint8
	GetHands() []hand.Hand

	GetCurrentHandCardsPointer() *hand.Hand
	GetCurrentHandCards() hand.Hand

	SetBalance(int)
	GetBalance() int

	ResetHands()
	SplitHands() error

	PlaceInsurance() error

	SetCurrentHandWinnings(outcome.Winnings) error
	SetCurrentHandOutcome(outcome.BlackjackOutcome) error

	GetCurrentHandWinnings() outcome.Winnings
	GetCurrentHandOutcome() outcome.BlackjackOutcome
}

func New(balance int) *player {
	p := player{
		balance:          balance,
		hands:            make([]playerHand, 1),
		totalHands:       1,
		currentHandIndex: 0,
	}
	return &p
}

func (player *player) getMoneyFromBalance(amount int) (int, error) {
	if amount <= player.GetBalance() {
		player.SetBalance(player.GetBalance() - amount)
		return amount, nil
	}
	return 0, fmt.Errorf(blackjackErrors.InvalidBalance)
}

func (player *player) PlaceInsurance() error {
	if player.totalHands == 1 && len(player.hands[0].cards) == 2 {
		insurance, err := player.getMoneyFromBalance(player.GetCurrentHandBet() / 2)
		if err == nil {
			player.hands[0].insurance = insurance
			return nil
		}
		return err
	}
	return fmt.Errorf(blackjackErrors.InvalidInsuranceHand)
}

func (player player) GetCurrentHandWinnings() outcome.Winnings {
	return player.getCurrentHand().winnings
}

func (player player) GetCurrentHandOutcome() outcome.BlackjackOutcome {
	return player.getCurrentHand().outcome
}

func (player player) GetHands() []hand.Hand {
	var hands []hand.Hand
	for _, hand := range player.hands {
		hands = append(hands, hand.cards)
	}
	return hands
}

func (player player) getCurrentHand() *playerHand {
	return &player.hands[player.GetCurrentHandIndex()]
}

func (player *player) SetCurrentHandWinnings(winnings outcome.Winnings) error {
	if player.GetCurrentHandBet() == 0 || len(player.GetCurrentHandCards()) < 2 {
		return fmt.Errorf(blackjackErrors.InvalidSetWinningsHand)
	}

	player.getCurrentHand().winnings = winnings
	return nil
}

func (player *player) SetCurrentHandOutcome(outcome outcome.BlackjackOutcome) error {
	if player.GetCurrentHandBet() == 0 || len(player.GetCurrentHandCards()) < 2 {
		return fmt.Errorf(blackjackErrors.InvalidSetOutcomeHand)
	}

	player.getCurrentHand().outcome = outcome
	return nil
}

func (player player) GetCurrentHandBet() int {
	return player.hands[player.currentHandIndex].bet
}

func (player *player) SetCurrentHandBet(bet int) error {
	if player.hands[player.currentHandIndex].bet != 0 {
		return fmt.Errorf(blackjackErrors.BetAlreadyPlaced)
	}
	if bet < 1 {
		return fmt.Errorf("bet has to be greater than 0")
	}

	if bet > player.balance {
		return fmt.Errorf(blackjackErrors.BetHigherThanBalance)
	}

	validBet, err := player.getMoneyFromBalance(bet)
	if err == nil {
		player.hands[player.currentHandIndex].bet = validBet
		return nil
	}
	return err
}

func (player *player) DoubleCurrentHandBet() error {
	if player.hands[player.currentHandIndex].doubleDown {
		return fmt.Errorf(blackjackErrors.BetAlreadyDoubled)
	}

	if player.hands[player.currentHandIndex].bet == 0 {
		return fmt.Errorf(blackjackErrors.NoBetsPlaced)
	}

	if len(player.GetCurrentHandCards()) != 2 {
		return fmt.Errorf(blackjackErrors.InvalidCardsForDoubleDown)
	}

	if player.GetBalance() < player.GetCurrentHandBet() {
		return fmt.Errorf(blackjackErrors.NoMoneyForDoubleDown)
	}

	if player.GetCurrentHandBet()*2 > player.balance {
		player.hands[player.currentHandIndex].bet += player.balance
		player.hands[player.currentHandIndex].doubleDown = true
		player.balance = 0

		return fmt.Errorf("bet given higher then balance. bet set to balance value")
	}
	player.balance -= player.GetCurrentHandBet()
	player.hands[player.currentHandIndex].bet += player.GetCurrentHandBet()
	player.hands[player.currentHandIndex].doubleDown = true
	return nil
}

func (player *player) SetCurrentHandIndex(currentHandIndex uint8) error {
	if currentHandIndex > player.totalHands-1 {
		return fmt.Errorf("hand index given bigger than total hands")
	}
	player.currentHandIndex = currentHandIndex
	return nil
}

func (player player) GetCurrentHandIndex() uint8 {
	return player.currentHandIndex
}

func (player *player) GetCurrentHandCardsPointer() *hand.Hand {
	return &player.hands[player.currentHandIndex].cards
}

func (player player) GetCurrentHandCards() hand.Hand {
	return player.hands[player.currentHandIndex].cards
}

func (player *player) SetBalance(balance int) {
	player.balance = balance
}

func (player *player) GetBalance() int {
	return player.balance
}

func (player *player) ResetHands() {
	player.hands = make([]playerHand, 1)
}

func (player player) GetTotalHands() uint8 {
	return player.totalHands
}

func (player *player) SplitHands() error {
	handToBeSplitted := player.GetCurrentHandCards()
	handToBeSplittedBet := player.GetCurrentHandBet()
	if len(handToBeSplitted) != 2 {
		return fmt.Errorf(blackjackErrors.SplitCardsNumberError)
	}

	if handToBeSplitted[0:1].Score() != handToBeSplitted[1:2].Score() {
		return fmt.Errorf(blackjackErrors.SplitCardsValueError)
	}

	if player.GetBalance() < player.GetCurrentHandBet() {
		return fmt.Errorf(blackjackErrors.NoMoneyForSplitting)
	}

	newHand := playerHand{}
	newHand.cards = hand.Hand{handToBeSplitted[1]}

	player.GetCurrentHandCardsPointer().RemoveCardAtIndex(1)

	player.hands = append(player.hands, newHand)
	player.totalHands++
	player.currentHandIndex++

	_ = player.SetCurrentHandBet(handToBeSplittedBet)

	return nil
}
