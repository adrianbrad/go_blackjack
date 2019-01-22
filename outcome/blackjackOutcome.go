package outcome

import (
	"blackjack/blackjackWinner"
	"blackjack/hand"
)

type BlackjackOutcome struct {
	Winner         blackjackWinner.BlackjackWinner
	theOtherBusted bool
	blackjack      bool
}

func NewBlackjackOutcome(winner blackjackWinner.BlackjackWinner, theOtherBusted, blackjack bool) BlackjackOutcome {
	return BlackjackOutcome{winner, theOtherBusted, blackjack}
}

func ComputeOutcome(playerHand hand.Hand, dealerHand hand.Hand) BlackjackOutcome {
	playerScore := playerHand.Score()
	dealerScore := dealerHand.Score()

	playerBlackjack := playerHand.Blackjack()
	dealerBlackjack := dealerHand.Blackjack()
	switch {
	case playerBlackjack && dealerBlackjack:
		return NewBlackjackOutcome(blackjackWinner.Draw, false, true)
	case dealerBlackjack:
		return NewBlackjackOutcome(blackjackWinner.Dealer, false, true)
	case playerBlackjack:
		return NewBlackjackOutcome(blackjackWinner.Player, false, true)
	case playerScore > 21: //if player busts nothing else matters
		return NewBlackjackOutcome(blackjackWinner.Dealer, true, false)
	case dealerScore > 21:
		return NewBlackjackOutcome(blackjackWinner.Player, true, false)
	case playerScore > dealerScore:
		return NewBlackjackOutcome(blackjackWinner.Player, false, false)
	case playerScore < dealerScore:
		return NewBlackjackOutcome(blackjackWinner.Dealer, false, false)
	case playerScore == dealerScore:
		return NewBlackjackOutcome(blackjackWinner.Draw, false, false)
	default:
		panic("Something's wrong with the outcome")
	}
}
