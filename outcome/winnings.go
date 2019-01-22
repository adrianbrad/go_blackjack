package outcome

import "blackjack/blackjackWinner"

type Winnings int
type BetBack int

func ComputeWinningsForPlayer(outcome BlackjackOutcome, playerBet int, blackjackPayout float64) (Winnings, BetBack) {
	switch outcome.Winner {
	case blackjackWinner.Player:
		if outcome.blackjack {
			return Winnings(float64(playerBet) * blackjackPayout), BetBack(playerBet)
		}
		return Winnings(playerBet), BetBack(playerBet)
	case blackjackWinner.Dealer:
		return Winnings(0), BetBack(0)
	case blackjackWinner.Draw:
		return Winnings(0), BetBack(playerBet)
	default:
		panic("ERROR: Wrong outcome")
	}
}
