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

	Dealer Dealer
}

func (game *Game) GetCurrentPlayerHand() *hand.Hand {
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
	*game.GetCurrentPlayerHand() = append(*game.GetCurrentPlayerHand(), game.Deck.DealCard())
}

func (game *Game) Stand() {
	switch game.State {
	case StatePlayerTurn:
		game.State = StateDealerTurn
	case StateDealerTurn:
		game.State = StateHandOver
	}
}

func (game *Game) DealStartingHands() { //FIXME bullshit
	game.State = StatePlayerTurn
	game.Hit()
	game.State = StateDealerTurn
	game.Hit()
	game.State = StatePlayerTurn
	game.Hit()
	game.State = StateDealerTurn
	game.Hit()

	game.State = StatePlayerTurn
}

func (game *Game) FinishDealerHand() {
	decision := game.Dealer.TakeDecision(game.DealerHand)
	decision(game)
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
	_, playerScore := playerHand.Score()
	_, dealerScore := dealerHand.Score()
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
	default:
		panic("Something's wrong with the outcome")
	}
}
