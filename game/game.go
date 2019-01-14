package game

import (
	"blackjack/blackjackWinner"
	"blackjack/hand"
	"blackjack/player"
	"deck"
	"fmt"
	"reflect"
)

type GameSessionState uint8

const (
	StateBet GameSessionState = iota
	StatePlayerTurn
	StateDealerTurn
	StateHandOver
)

type Game struct {
	NumDecks        int
	BlackjackPayout float64

	Deck  deck.Deck
	State GameSessionState

	Player        player.Player
	CurrentPlayer uint8
	TotalPlayers  uint8

	DealerHand hand.Hand
	Dealer     Dealer
}

func (game *Game) Bet(bet int) {
	if game.State != StateBet {

	}
	_ = game.Player.SetCurrentHandBet(bet)
}

func (game *Game) DoubleDown() { //TODO: handle len(game.PlayerHand) > 2
	if game.State != StatePlayerTurn {

	}

	if len(game.Player.GetCurrentHandCards()) != 2 {
		//handle
	}
	_ = game.Player.DoubleCurrentHandBet()
	game.Hit()
	game.Stand()
}

func (game *Game) Split() {
	if game.State != StatePlayerTurn {

	}
	if len(game.Player.GetCurrentHandCards()) != 2 {
		//you can only split with two cards in hand
	}
	if game.Player.GetCurrentHandCards()[0].Rank != game.Player.GetCurrentHandCards()[0].Rank {
		//you can only split cards with same rank
	}
}

func (game *Game) GetCurrentPlayerHand() *hand.Hand {
	switch game.State {
	case StatePlayerTurn:
		return game.Player.GetCurrentHandCardsPointer()
	case StateDealerTurn:
		return &game.DealerHand
	default:
		panic("ERROR: it isn't currently any player turn")
	}
}

func (game *Game) ShuffleNewDeck() {
	game.Deck = deck.New(deck.Amount(game.NumDecks), deck.Shuffle)
}

func (game *Game) Hit() {
	game.GetCurrentPlayerHand().AddCard(game.Deck.DealCard())
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
	for i := 0; i < 2; i++ {
		game.State = StatePlayerTurn
		game.Hit()
		game.State = StateDealerTurn
		game.Hit()
	}
	//game.PlayerHand = hand.Hand{{deck.Clubs,deck.Six}, {deck.Clubs, deck.Five}, {deck.Clubs, deck.Ten}}
	//game.DealerHand = hand.Hand{{deck.Clubs,deck.Ace}, {deck.Clubs, deck.King}}

	game.State = StatePlayerTurn
}

func (game *Game) FinishDealerHand() {
	if game.State != StateDealerTurn {

	}
	decision := game.Dealer.TakeDecision(game.DealerHand)

	for reflect.ValueOf(decision).Pointer() == reflect.ValueOf((*Game).Hit).Pointer() { //while the dealer decides to hit, execute the hit method, crazy stuff i know
		decision(game)
		decision = game.Dealer.TakeDecision(game.DealerHand)
	}
	decision(game) //the dealer stands here
}

func (game *Game) EndHand() { //TODO

	outcome := computeOutcome(game.Player.GetCurrentHandCards(), game.DealerHand)
	winnings := game.computeWinningsForPlayer(outcome, game.Player.GetCurrentHandBet())
	moneyOperations := game.computeMoneyOperations(winnings)

	fmt.Println(moneyOperations)

	game.Player.SetBalance(game.Player.GetBalance() + int(winnings))

	game.Player.ResetHands()
	game.DealerHand = nil
	fmt.Println(outcome)
	fmt.Println(winnings)
	fmt.Println(game.Player.GetBalance())

	min := 52 * game.NumDecks / 3 //reshuffle after we consumed 2/3
	if len(game.Deck) < min {
		game.ShuffleNewDeck()
	}
}

func (game *Game) computeMoneyOperations(w winnings) []MoneyOperation { //TODO

	return []MoneyOperation{{betBack, game.Player.GetCurrentHandBet()}}
}

type winnings int

func (game *Game) computeWinningsForPlayer(outcome BlackjackOutcome, playerBet int) winnings {
	switch outcome.winner {
	case blackjackWinner.Player:
		if outcome.blackjack {
			return winnings(int(float64(playerBet) * game.BlackjackPayout))
		}
		return winnings(playerBet * 2)
	case blackjackWinner.Dealer:
		return winnings(0)
	case blackjackWinner.Draw:
		return winnings(playerBet)
	default:
		panic("ERROR: Wrong outcome")
	}
}

type MoneyOperationType uint8

type MoneyOperation struct {
	operationType MoneyOperationType
	amount        int
}

const (
	bet MoneyOperationType = iota
	betBack
	win
)

type BlackjackOutcome struct {
	winner         blackjackWinner.BlackjackWinner
	theOtherBusted bool
	blackjack      bool
}

func computeOutcome(playerHand hand.Hand, dealerHand hand.Hand) BlackjackOutcome {
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
