package hand

import "deck"

type Hand []deck.Card

func (hand Hand) minScore() int {
	score := 0
	for _, card := range hand {
		score += min(int(card.Rank), 10)
	}
	return score
}

func min(number, max int) int {
	if number < max {
		return number
	}
	return max
}

func (hand Hand) Score() int {

	//ace is counted as a one if the score is higher than 11
	minScore := hand.minScore()
	if minScore > 11 {
		return minScore
	}

	//ace is counted as 11 if the score is lower than 11
	for _, card := range hand {
		if card.Rank == deck.Ace {

			//adding ten as we already have a +1 if an ace is in the hand
			return minScore + 10
		}
	}

	//if no ace is in the hand
	return minScore
}

func DealCard(deck *deck.Deck, hand *Hand) {
	*hand, *deck = append(*hand, (*deck)[0]), (*deck)[1:]
}
