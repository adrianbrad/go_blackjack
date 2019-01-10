package ai

import (
	"blackjack/game"
	"blackjack/hand"
	"deck"
)

type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Decision
}

type HumanAI struct {
}

func (ai *HumanAI) Play(playerHand hand.Hand, dealerVisibleCard deck.Card) Decision { //[]hand.Hand - a slice of hand.Hand(which is also a slice) - in case of splitting user will have multiple hands
	return func(i game.Game) game.Game {
		return game.Game{}
	}
}

func (ai *HumanAI) Bet() uint8 {
	return 1
}

type Decision func(game.Game) game.Game
