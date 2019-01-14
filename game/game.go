package game

import (
	"blackjack/dealer"
	"blackjack/hand"
	"blackjack/outcome"
	"blackjack/player"
	"deck"
	"fmt"
)

type GameSessionState uint8

const (
	StateBet GameSessionState = iota
	StatePlayerTurn
	StateDealerTurn
	StateHandOver
)

type Game interface {
	DealStartingHands()
	Bet(int)
	PlaceInsurance()
	DoubleDown()
	Split()
	ShuffleNewDeck()
	Hit()
	Stand()
	FinishDealerHand()
	EndHand()

	GetPlayer() player.Player
	GetDealer() dealer.Dealer
	GetDeck() deck.Deck
	GetState() GameSessionState
	GetBlackjackPayout() float64
}

type game struct {
	numDecks        int
	blackjackPayout float64

	deck  deck.Deck
	state GameSessionState

	player        player.Player
	currentPlayer uint8
	totalPlayers  uint8

	dealer dealer.Dealer
}

func New(numDecks int, blackjackPayout float64, player player.Player, dealer dealer.Dealer) *game {
	return &game{
		numDecks:        numDecks,
		blackjackPayout: blackjackPayout,
		player:          player,
		dealer:          dealer,
	}
}

func (game *game) Bet(bet int) {
	if game.state != StateBet {

	}
	_ = game.player.SetCurrentHandBet(bet)
}

func (game *game) PlaceInsurance() {

}

func (game *game) DoubleDown() {
	if game.state != StatePlayerTurn {

	}
	_ = game.player.DoubleCurrentHandBet()
	game.Hit()
	game.Stand()
}

func (game *game) Split() {
	if game.state != StatePlayerTurn {

	}
	if len(game.player.GetCurrentHandCards()) != 2 {
		//you can only split with two cards in hand
	}
	if game.player.GetCurrentHandCards()[0].Rank != game.player.GetCurrentHandCards()[0].Rank {
		//you can only split cards with same rank
	}

	err := game.player.SplitHands()
	fmt.Println(err)
}

func (game *game) getCurrentPlayerHand() *hand.Hand {
	switch game.state {
	case StatePlayerTurn:
		return game.player.GetCurrentHandCardsPointer()
	case StateDealerTurn:
		return game.dealer.GetDealerHandPointer()
	default:
		panic("ERROR: it isn't currently any player turn")
	}
}

func (game *game) ShuffleNewDeck() {
	game.deck = deck.New(deck.Amount(game.numDecks), deck.Shuffle)
}

func (game *game) Hit() {
	game.getCurrentPlayerHand().AddCard(game.deck.DealCard())
}

func (game *game) Stand() {
	switch game.state {
	case StatePlayerTurn:
		game.state = StateDealerTurn
	case StateDealerTurn:
		game.state = StateHandOver
	}
}

func (game *game) DealStartingHands() { //FIXME bullshit changing states
	for i := 0; i < 2; i++ {
		game.state = StatePlayerTurn
		game.Hit()
		game.state = StateDealerTurn
		game.Hit()
	}
	//game.PlayerHand = hand.Hand{{deck.Clubs,deck.Six}, {deck.Clubs, deck.Five}, {deck.Clubs, deck.Ten}}
	//game.DealerHand = hand.Hand{{deck.Clubs,deck.Ace}, {deck.Clubs, deck.King}}

	game.state = StatePlayerTurn
}

func (game *game) FinishDealerHand() {
	if game.state != StateDealerTurn {

	}

	canHit := game.dealer.CanHit()

	for canHit { //while the dealer decides to hit, execute the hit method, crazy stuff i know
		game.Hit()
		canHit = game.dealer.CanHit()
	}

	game.Stand()
}

func (game *game) EndHand() { //TODO

	result := outcome.ComputeOutcome(game.player.GetCurrentHandCards(), game.dealer.GetDealerHand())
	winnings := outcome.ComputeWinningsForPlayer(result, game.player.GetCurrentHandBet(), game.GetBlackjackPayout())
	moneyOperations := outcome.ComputeMoneyOperations(winnings, game.GetPlayer().GetCurrentHandBet())

	fmt.Println(moneyOperations)

	game.player.SetBalance(game.player.GetBalance() + int(winnings))

	game.player.ResetHands()
	game.dealer.ResetHands()

	fmt.Println(result)
	fmt.Println(winnings)
	fmt.Println(game.player.GetBalance())

	min := 52 * game.numDecks / 3 //reshuffle after we consumed 2/3
	if len(game.deck) < min {
		game.ShuffleNewDeck()
	}
}

func (game game) GetDealer() dealer.Dealer {
	return game.dealer
}

func (game game) GetDeck() deck.Deck {
	return game.deck
}

func (game game) GetPlayer() player.Player {
	return game.player
}

func (game game) GetState() GameSessionState {
	return game.state
}

func (game game) GetBlackjackPayout() float64 {
	return game.blackjackPayout
}
