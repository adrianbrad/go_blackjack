package game

import "blackjack/hand"

type Dealer interface {
	TakeDecision(hand hand.Hand) Decision
}

type BasicDealer struct{}

func (dealer BasicDealer) TakeDecision(hand hand.Hand) Decision {
	_, dealerScore := hand.Score()
	if dealerScore <= 16 || (dealerScore == 17 && hand.Soft()) { //(hand.Score() == 17 && hand.MinScore() != 17) - this means it has an ace, and it's a soft 17
		return (*Game).Hit
	} else {
		return (*Game).Stand
	}
}

type Decision func(*Game)
