package dealer

import (
	"blackjack/hand"
	"fmt"
	"github.com/adrianbrad/go-deck-of-cards"
)

type Dealer interface {
	CanHit() bool
	GetDealerHand() hand.Hand
	GetDealerHandPointer() *hand.Hand
	GetDealerFirstCard() (*deck.Card, error)
	ResetHands()
}

type defaultDealer struct {
	hand hand.Hand
}

func NewDefaultDealer() *defaultDealer {
	return &defaultDealer{}
}

func (dealer defaultDealer) CanHit() bool { //we can just put (defaultDealer) instead of (dealer defaultDealer) as we store no state in Basic dealer
	dealerScore := dealer.GetDealerHand().Score()
	if dealerScore <= 16 || (dealerScore == 17 && dealer.GetDealerHand().Soft()) { //(hand.Score() == 17 && hand.MinScore() != 17) - this means it has an ace, and it's a soft 17
		return true
	}
	return false
}

func (dealer *defaultDealer) GetDealerHandPointer() *hand.Hand {
	return &dealer.hand
}

func (dealer defaultDealer) GetDealerHand() hand.Hand {
	return dealer.hand
}

func (dealer *defaultDealer) ResetHands() {
	dealer.hand = hand.Hand{}
}

func (dealer defaultDealer) GetDealerFirstCard() (*deck.Card, error) {
	if len(dealer.hand) < 1 {
		return nil, fmt.Errorf("dealer has no cards")
	}
	return &dealer.GetDealerHand()[0], nil
}
