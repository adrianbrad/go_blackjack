package outcome

import (
	"blackjack/blackjackWinner"
	"blackjack/hand"
)

type BlackjackOutcome struct {
	winner         blackjackWinner.BlackjackWinner
	theOtherBusted bool
	blackjack      bool
}

func ComputeOutcome(playerHand hand.Hand, dealerHand hand.Hand) BlackjackOutcome {
	playerScore := playerHand.Score()
	dealerScore := dealerHand.Score()

	playerBlackjack := playerHand.Blackjack()
	dealerBlackjack := dealerHand.Blackjack()
	switch {
	case playerBlackjack && dealerBlackjack:
		return BlackjackOutcome{blackjackWinner.Draw, false, true}
	case dealerBlackjack:
		return BlackjackOutcome{blackjackWinner.Dealer, false, true}
	case playerBlackjack:
		return BlackjackOutcome{blackjackWinner.Player, false, true}
	case playerScore > 21: //if player busts nothing else matters
		return BlackjackOutcome{blackjackWinner.Dealer, true, false}
	case dealerScore > 21:
		return BlackjackOutcome{blackjackWinner.Player, true, false}
	case playerScore > dealerScore:
		return BlackjackOutcome{blackjackWinner.Player, false, false}
	case playerScore < dealerScore:
		return BlackjackOutcome{blackjackWinner.Dealer, false, false}
	case playerScore == dealerScore:
		return BlackjackOutcome{blackjackWinner.Draw, false, false}
	default:
		panic("Something's wrong with the outcome")
	}
}
