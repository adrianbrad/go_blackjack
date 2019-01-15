package player

import (
	"blackjack/blackjackErrors"
	"blackjack/hand"
	"fmt"
)

type playerHand struct {
	hand       hand.Hand
	handBet    int
	betPlaced  bool
	doubledBet bool
}

type Player interface {
	SetCurrentHandBet(int) error
	GetCurrentHandBet() int
	DoubleCurrentHandBet() error

	SetCurrentIndexHand(uint8) error
	GetCurrentHandIndex() uint8

	GetTotalHands() uint8

	GetCurrentHandCardsPointer() *hand.Hand
	GetCurrentHandCards() hand.Hand

	SetBalance(int)
	GetBalance() int

	ResetHands()
	SplitHands() error
}

type player struct {
	hands            []playerHand
	currentHandIndex uint8
	totalHands       uint8
	balance          int
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

func (player player) GetCurrentHandBet() int {
	return player.hands[player.currentHandIndex].handBet
}

func (player *player) SetCurrentHandBet(bet int) error {
	if player.hands[player.currentHandIndex].betPlaced {
		return fmt.Errorf(blackjackErrors.BetAlreadyPlaced)
	}
	if bet < 1 {
		return fmt.Errorf("bet has to be greater than 0")
	}

	if bet > player.balance {
		return fmt.Errorf(blackjackErrors.BetHigherThanBalance)
	}
	player.balance -= bet
	player.hands[player.currentHandIndex].handBet = bet
	player.hands[player.currentHandIndex].betPlaced = true
	return nil
}

func (player *player) DoubleCurrentHandBet() error {
	if player.hands[player.currentHandIndex].doubledBet {
		return fmt.Errorf("bet already doubled")
	}

	if !player.hands[player.currentHandIndex].betPlaced {
		return fmt.Errorf("no bet placed")
	}

	if len(player.hands[player.currentHandIndex].hand) != 2 {
		return fmt.Errorf("current hand has more than 2 cards")
	}

	if player.GetCurrentHandBet()*2 > player.balance {
		player.hands[player.currentHandIndex].handBet += player.balance
		player.hands[player.currentHandIndex].doubledBet = true
		player.balance = 0

		return fmt.Errorf("bet given higher then balance. bet set to balance value")
	}
	player.balance -= player.GetCurrentHandBet()
	player.hands[player.currentHandIndex].handBet += player.GetCurrentHandBet()
	player.hands[player.currentHandIndex].doubledBet = true
	return nil
}

func (player *player) SetCurrentIndexHand(currentHandIndex uint8) error {
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
	return &player.hands[player.currentHandIndex].hand
}

func (player player) GetCurrentHandCards() hand.Hand {
	return player.hands[player.currentHandIndex].hand
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
		return fmt.Errorf("hand must have exactly two cards to split")
	}

	if handToBeSplitted[0].Rank != handToBeSplitted[1].Rank {
		return fmt.Errorf("the cards should have the same value for splitting")
	}

	newHand := playerHand{}
	newHand.hand = hand.Hand{handToBeSplitted[1]}

	player.GetCurrentHandCardsPointer().RemoveCardAtIndex(1)

	player.hands = append(player.hands, newHand)
	player.totalHands++
	player.currentHandIndex++

	_ = player.SetCurrentHandBet(handToBeSplittedBet)

	return nil
}
