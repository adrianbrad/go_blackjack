package hand

import (
	"deck"
)

type Hand []deck.Card

func (hand Hand) MinScore() ([]int, int) {
	score := 0
	var individualValues []int

	for _, card := range hand {
		cardValue := min(int(card.Rank), 10)
		score += cardValue

		individualValues = append(individualValues, cardValue)
	}
	return individualValues, score
}

func min(number, max int) int {
	if number < max {
		return number
	}
	return max
}

func (hand Hand) Score() ([]int, int) {

	//ace is counted as a one if the score is higher than 11
	individualValues, minScore := hand.MinScore()
	if minScore > 11 {
		return individualValues, minScore
	}

	//ace is counted as 11 if the score is lower than 11
	for index, card := range hand {
		if card.Rank == deck.Ace {
			individualValues[index] += 10
			//adding ten as we already have a +1 if an ace is in the hand
			return individualValues, minScore + 10
		}
	}

	//if no ace is in the hand
	return individualValues, minScore
}

func (hand Hand) Soft() bool {
	_, minScore := hand.MinScore()
	_, score := hand.Score()
	return minScore != score
}

//func DrawCard(deck *deck.Deck, hand *Hand) {
//	*hand = append(*hand, deck.DealCard())
//}
