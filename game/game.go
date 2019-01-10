package game

import (
	"blackjack/hand"
	"deck"
	"fmt"
)

type GameSessionState uint8

const (
	StatePlayerTurn GameSessionState = iota
	StateDealerTurn
	StateHandOver
)

type Game struct {
	Deck       deck.Deck
	State      GameSessionState //change to custom type
	PlayerHand hand.Hand
	DealerHand hand.Hand
}

func (game *Game) CurrentPlayerHand() *hand.Hand {
	switch game.State {
	case StatePlayerTurn:
		return &game.PlayerHand
	case StateDealerTurn:
		return &game.DealerHand
	default:
		panic("ERROR: it isn't currently any player turn")
	}
}

func (game *Game) ShuffleNewDeck() {
	game.Deck = deck.New(deck.Amount(3), deck.Shuffle)
}

func (game *Game) Hit() {
	hand.DealCard(&game.Deck, game.CurrentPlayerHand())
}

func (game *Game) Stand() {
	switch game.State {
	case StatePlayerTurn:
		game.State = StateDealerTurn
	case StateDealerTurn:
		game.State = StateHandOver
	}
}

func (game *Game) DealStartingHands() {
	hand.DealCard(&game.Deck, &game.PlayerHand)
	hand.DealCard(&game.Deck, &game.DealerHand)
	hand.DealCard(&game.Deck, &game.PlayerHand)
	hand.DealCard(&game.Deck, &game.DealerHand)
	game.State = StatePlayerTurn
}

func (game *Game) FinishDealerHand() {
	for game.DealerHand.Score() <= 16 || (game.DealerHand.Score() == 17 && game.DealerHand.Score() != 17) { //(hand.Score() == 17 && hand.MinScore() != 17) - this means it has an ace, and it's a soft 17
		hand.DealCard(&game.Deck, &game.DealerHand)
	}
	game.Stand() //StateDealerHand -> StateHandOver
}

func (game *Game) EndHand() {
	outcome := computeOutcome(game.PlayerHand, game.DealerHand)
	game.PlayerHand = nil
	game.DealerHand = nil
	fmt.Println(outcome)
}

type BlackjackWinner uint8

const (
	player BlackjackWinner = iota
	dealer
	draw
)

type BlackjackOutcome struct {
	winner         BlackjackWinner
	theOtherBusted bool
}

func computeOutcome(playerHand hand.Hand, dealerHand hand.Hand) BlackjackOutcome {
	playerScore := playerHand.Score()
	dealerScore := dealerHand.Score()
	switch {
	case playerScore > 21: //if player busts nothing else matters
		return BlackjackOutcome{dealer, true}
	case dealerScore > 21:
		return BlackjackOutcome{player, true}
	case playerScore > dealerScore:
		return BlackjackOutcome{player, false}
	case playerScore < dealerScore:
		return BlackjackOutcome{dealer, false}
	case playerScore == dealerScore:
		return BlackjackOutcome{draw, false}
	}
	return BlackjackOutcome{draw, false}
}
