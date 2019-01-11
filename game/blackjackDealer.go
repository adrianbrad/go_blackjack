package game

import "blackjack/hand"

type Dealer interface {
	TakeDecision(hand hand.Hand) DealerDecision
}

type BasicDealer struct{}

func (BasicDealer) TakeDecision(hand hand.Hand) DealerDecision { //we can just put (BasicDealer) instead of (dealer BasicDealer) as we store no state in Basic Dealer
	dealerScore := hand.Score()
	if dealerScore <= 16 || (dealerScore == 17 && hand.Soft()) { //(hand.Score() == 17 && hand.MinScore() != 17) - this means it has an ace, and it's a soft 17
		return (*Game).Hit
	}
	return (*Game).Stand
}

type DealerDecision func(*Game)
