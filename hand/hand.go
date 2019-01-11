package hand

import (
	"deck"
)

type Hand []deck.Card

func (hand Hand) MinScore() int {
	score := 0
	var individualValues []int

	for _, card := range hand {
		cardValue := min(int(card.Rank), 10)
		score += cardValue

		individualValues = append(individualValues, cardValue)
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
	minScore := hand.MinScore()
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

func (hand Hand) Soft() bool {
	minScore := hand.MinScore()
	score := hand.Score()
	return minScore != score
}

func (hand Hand) IndividualScore() []int { //this can be improved
	var individualValues []int

	for _, card := range hand {
		cardValue := min(int(card.Rank), 10)
		individualValues = append(individualValues, cardValue)
	}

	for index, card := range hand {
		if card.Rank == deck.Ace {
			individualValues[index] += 10
			break
		}
	}
	return individualValues
}

func (hand Hand) Blackjack() bool { //Returns true if hand is blackjack
	return hand.Score() == 21 && len(hand) == 2
}

//func DrawCard(deck *deck.Deck, hand *Hand) {
//	*hand = append(*hand, deck.DealCard())
//}
