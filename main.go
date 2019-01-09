package main

import (
	"deck"
	"fmt"
)

type Hand []deck.Card

func (hand Hand) MinScore() int {
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

func main() {
	cards := deck.New(deck.Amount(3), deck.Shuffle)
	var player, dealer Hand
	dealCard(&cards, &dealer)
	dealCard(&cards, &player)
	dealCard(&cards, &dealer)
	dealCard(&cards, &player)
	fmt.Println(dealer.Score())
	finishDealerHand(&cards, &dealer)
	fmt.Println(player)
	fmt.Println(dealer)
	fmt.Println(player.Score())
	fmt.Println(dealer.Score())
	computeOutcome(player, dealer)
}

func finishDealerHand(deck *deck.Deck, hand *Hand) {
	for hand.Score() <= 16 || (hand.Score() == 17 && hand.MinScore() != 17) {
		dealCard(deck, hand)
	}
}

func dealCard(deck *deck.Deck, hand *Hand) {
	*hand, *deck = append(*hand, (*deck)[0]), (*deck)[1:]
}

func computeOutcome(playerHand Hand, dealerHand Hand) {
	playerScore := playerHand.Score()
	dealerScore := dealerHand.Score()
	switch {
	case playerScore > 21: //if player busts nothing else matters
		fmt.Println("Player busted")
	case dealerScore > 21:
		fmt.Println("Dealer busted")
	case playerScore > dealerScore:
		fmt.Println("Player win")
	case playerScore < dealerScore:
		fmt.Println("Player lose")
	case playerScore == dealerScore:
		fmt.Println("Draw")
	}
}
