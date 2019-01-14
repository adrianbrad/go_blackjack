package outcome

import "blackjack/blackjackWinner"

type Winnings int

func ComputeWinningsForPlayer(outcome BlackjackOutcome, playerBet int, blackjackPayout float64) Winnings {
	switch outcome.winner {
	case blackjackWinner.Player:
		if outcome.blackjack {
			return Winnings(int(float64(playerBet) * blackjackPayout))
		}
		return Winnings(playerBet * 2)
	case blackjackWinner.Dealer:
		return Winnings(0)
	case blackjackWinner.Draw:
		return Winnings(playerBet)
	default:
		panic("ERROR: Wrong outcome")
	}
}
