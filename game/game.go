package game

import (
	"blackjack/blackjackErrors"
	"blackjack/dealer"
	"blackjack/gameSessionState"
	"blackjack/hand"
	"blackjack/outcome"
	"blackjack/player"
	"fmt"
	"github.com/adrianbrad/go-deck-of-cards"
)

type Game interface {
	DealStartingHands() error
	Bet(int) error
	PlaceInsurance() error
	DoubleDown() error
	Split() error
	ShuffleNewDeck()
	Hit() error
	Stand() error
	FinishDealerHand() error
	EndHand() (outcome.BlackjackOutcome, outcome.Winnings, []outcome.MoneyOperation, error)

	GetPlayer() player.Player
	GetDealer() dealer.Dealer
	GetDeck() deck.Deck
	GetState() gameSessionState.GameSessionState
	GetBlackjackPayout() float64
}

type game struct {
	numDecks        int
	blackjackPayout float64

	deck        deck.Deck
	initialDeck deck.Deck
	state       gameSessionState.GameSessionState

	player        player.Player
	currentPlayer uint8
	totalPlayers  uint8

	dealer dealer.Dealer
}

func New(numDecks int, blackjackPayout float64, player player.Player, dealer dealer.Dealer, deck deck.Deck) *game {
	g := game{
		numDecks:        numDecks,
		blackjackPayout: blackjackPayout,
		player:          player,
		dealer:          dealer,
		deck:            deck,
	}
	if deck == nil {
		g.ShuffleNewDeck()
	}
	g.initialDeck = g.GetDeck()
	return &g

}

func (game *game) Bet(bet int) error {
	err := game.checkValidState(gameSessionState.StateBet)
	if err != nil {
		return err
	}

	err = game.player.SetCurrentHandBet(bet)
	if err != nil {
		return err
	}
	return nil
}

func (game *game) PlaceInsurance() error {
	err := game.checkValidState(gameSessionState.StatePlayerTurn)
	if err != nil {
		return err
	}

	dealerFirstCard, _ := game.GetDealer().GetDealerFirstCard()

	if dealerFirstCard.Score() != 11 {
		return fmt.Errorf(blackjackErrors.DealerFirstCardError)
	}

	game.GetPlayer().PlaceInsurance()

	if game.GetDealer().GetDealerHand().Blackjack() {
		game.state = gameSessionState.StateHandOver
	}

	return nil
}

func (game *game) DoubleDown() error {
	err := game.checkValidState(gameSessionState.StatePlayerTurn)
	if err != nil {
		return err
	}

	err = game.player.DoubleCurrentHandBet()
	if err != nil {
		return err
	}

	_ = game.Hit()
	game.Stand()

	return nil
}

func (game *game) Split() error {
	err := game.checkValidState(gameSessionState.StatePlayerTurn)
	if err != nil {
		return err
	}

	err = game.player.SplitHands()

	return err
}

func (game *game) getCurrentPlayerHand() (*hand.Hand, error) {
	switch game.state {
	case gameSessionState.StatePlayerTurn:
		return game.player.GetCurrentHandCardsPointer(), nil
	case gameSessionState.StateDealerTurn:
		return game.dealer.GetDealerHandPointer(), nil
	default:
		return nil, fmt.Errorf(blackjackErrors.NoActiveHands)
	}
}

func (game *game) ShuffleNewDeck() {
	game.deck = deck.New(deck.Amount(game.numDecks), deck.Shuffle)
}

func (game *game) Hit() error { //game can end from a hit that busts
	currentPlayerHand, err := game.getCurrentPlayerHand()
	if err != nil {
		return err
	}

	currentPlayerHand.AddCard(game.deck.DealCard())

	if currentPlayerHand.Score() > 21 { //if player busts go directly to endgame
		game.Stand()
	}

	return nil
}

func (game *game) Stand() error {

	switch game.state {
	case gameSessionState.StatePlayerTurn:
		game.state = gameSessionState.StateDealerTurn
		game.FinishDealerHand()
	case gameSessionState.StateDealerTurn:
		game.state = gameSessionState.StateHandOver
	default:
		return fmt.Errorf(blackjackErrors.NoActiveHands)
	}

	return nil
}

func (game *game) DealStartingHands() error { //FIXME bullshit changing states
	err := game.checkValidState(gameSessionState.StateBet)
	if err != nil {
		return err
	}
	if game.GetPlayer().GetCurrentHandBet() <= 0 {
		return fmt.Errorf("no bets placed")
	}

	for i := 0; i < 2; i++ {
		game.state = gameSessionState.StatePlayerTurn
		_ = game.Hit()
		game.state = gameSessionState.StateDealerTurn
		_ = game.Hit()
	}

	game.state = gameSessionState.StatePlayerTurn
	return nil
}

func (game *game) FinishDealerHand() error {
	err := game.checkValidState(gameSessionState.StateDealerTurn)
	if err != nil {
		return err
	}

	canHit := game.dealer.CanHit()

	for canHit { //while the dealer decides to hit, execute the hit method, crazy stuff i know
		_ = game.Hit()
		canHit = game.dealer.CanHit()
	}

	if game.state == gameSessionState.StateDealerTurn { //dealer can bust and we get to StateEndHand
		game.Stand()
	}

	return nil
}

func (game *game) EndHand() (outcome.BlackjackOutcome, outcome.Winnings, []outcome.MoneyOperation, error) { //TODO resolve insurance

	err := game.checkValidState(gameSessionState.StateHandOver)
	if err != nil {
		return outcome.BlackjackOutcome{}, outcome.Winnings(0), []outcome.MoneyOperation{}, err
	}

	result := outcome.ComputeOutcome(game.player.GetCurrentHandCards(), game.dealer.GetDealerHand())
	winnings := outcome.ComputeWinningsForPlayer(result, game.player.GetCurrentHandBet(), game.GetBlackjackPayout())
	moneyOperations := outcome.ComputeMoneyOperations(winnings, game.GetPlayer().GetCurrentHandBet())

	game.player.SetBalance(game.player.GetBalance() + int(winnings))

	game.player.ResetHands()
	game.dealer.ResetHands()
	game.state = gameSessionState.StateBet

	min := 52 * game.numDecks / 3 //reshuffle after we consumed 2/3
	if len(game.deck) < min {
		game.ShuffleNewDeck()
	}

	return result, winnings, moneyOperations, nil
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

func (game game) GetState() gameSessionState.GameSessionState {
	return game.state
}

func (game game) GetBlackjackPayout() float64 {
	return game.blackjackPayout
}

func (game game) checkValidState(givenState gameSessionState.GameSessionState) error {
	if game.GetState() != givenState {
		return fmt.Errorf("given state: %s different from the current state: %s", givenState, game.GetState())
	}
	return nil
}
