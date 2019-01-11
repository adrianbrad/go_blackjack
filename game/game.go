package game

import (
	"blackjack/hand"
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
	State GameSessionState //change to custom type

	PlayerHand    hand.Hand
	PlayerBet     int
	PlayerBalance int

	DealerHand hand.Hand
	Dealer     Dealer
}

func (game *Game) Bet(bet int) {
	game.PlayerBalance -= bet
	game.PlayerBet = bet
}

func (game *Game) DoubleDown() { //TODO: handle len(game.PlayerHand) > 2
	if game.State != StatePlayerTurn {

	}
	if len(game.PlayerHand) != 2 {
		//handle
	}
	game.PlayerBet *= 2
	game.Hit()
	game.Stand()
}

func (game *Game) Split() {
	if game.State != StatePlayerTurn {

	}
	if len(game.PlayerHand) != 2 {
		//you can only split with two cards in hand
	}
	if game.PlayerHand[0].Rank != game.PlayerHand[0].Rank {
		//you can only split cards with same rank
	}
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
	game.Deck = deck.New(deck.Amount(game.NumDecks), deck.Shuffle)
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
	decision := game.Dealer.TakeDecision(game.DealerHand)

	for reflect.ValueOf(decision).Pointer() == reflect.ValueOf((*Game).Hit).Pointer() { //while the dealer decides to hit, crazy stuff i know
		decision(game)
		decision = game.Dealer.TakeDecision(game.DealerHand)
	}
	decision(game) //the dealer stands here
}

func (game *Game) EndHand() { //TODO

	outcome := computeOutcome(game.PlayerHand, game.DealerHand)
	winnings := game.computeWinningsForPlayer(outcome, game.PlayerBet)
	moneyOperations := game.computeMoneyOperations(winnings)

	fmt.Println(moneyOperations)

	game.PlayerBalance += int(winnings)

	game.PlayerHand = nil
	game.DealerHand = nil
	fmt.Println(outcome)
	fmt.Println(winnings)
	fmt.Println(game.PlayerBalance)

	min := 52 * game.NumDecks / 3 //reshuffle after we consumed 2/3
	if len(game.Deck) < min {
		game.ShuffleNewDeck()
	}
}

func (game *Game) computeMoneyOperations(w winnings) []MoneyOperation { //TODO

	return []MoneyOperation{{betBack, game.PlayerBet}}
}

type winnings int

func (game *Game) computeWinningsForPlayer(outcome BlackjackOutcome, playerBet int) winnings {
	switch outcome.winner {
	case player:
		if outcome.blackjack {
			return winnings(int(float64(playerBet) * game.BlackjackPayout))
		}
		return winnings(playerBet * 2)
	case dealer:
		return winnings(0)
	case draw:
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

type BlackjackWinner uint8

const (
	player BlackjackWinner = iota
	dealer
	draw
)

type BlackjackOutcome struct {
	winner         BlackjackWinner
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
		return BlackjackOutcome{draw, false, true}
	case dealerBlackjack:
		return BlackjackOutcome{dealer, false, true}
	case playerBlackjack:
		return BlackjackOutcome{player, false, true}
	case playerScore > 21: //if player busts nothing else matters
		return BlackjackOutcome{dealer, true, false}
	case dealerScore > 21:
		return BlackjackOutcome{player, true, false}
	case playerScore > dealerScore:
		return BlackjackOutcome{player, false, false}
	case playerScore < dealerScore:
		return BlackjackOutcome{dealer, false, false}
	case playerScore == dealerScore:
		return BlackjackOutcome{draw, false, false}
	default:
		panic("Something's wrong with the outcome")
	}
}
