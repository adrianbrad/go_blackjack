package game

import (
	"blackjack/blackjackErrors"
	"blackjack/dealer"
	"blackjack/gameSessionState"
	"blackjack/hand"
	"blackjack/outcome"
	"blackjack/player"
	"fmt"

	deck "github.com/adrianbrad/go-deck-of-cards"
)

//Game defines
type Game interface {
	Bet(int) error
	DealStartingHands() error

	Hit() error
	Stand() error
	PlaceInsurance() error
	DoubleDown() error
	Split() error

	FinishDealerHand() error
	EndHand() ([]outcome.BlackjackOutcome, outcome.Winnings, [][]outcome.MoneyOperation, error)

	ShuffleNewDeck()

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

//New creates
func New(numDecks int, blackjackPayout float64, player player.Player, dealer dealer.Dealer, deck deck.Deck) Game {
	var g Game
	g = &game{
		numDecks:        numDecks,
		blackjackPayout: blackjackPayout,
		player:          player,
		dealer:          dealer,
		deck:            deck,
	}
	if deck == nil {
		g.ShuffleNewDeck()
	}

	return g
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

func (game *game) DealStartingHands() error {
	err := game.checkValidState(gameSessionState.StateBet)
	if err != nil {
		return err
	}

	game.state = gameSessionState.StatePlayerTurn
	if game.GetPlayer().GetCurrentHandBet() <= 0 {
		return fmt.Errorf("no bets placed")
	}

	for i := 0; i < 2; i++ {
		_ = game.Hit()
		if game.GetState() == gameSessionState.StateHandOver { //in case the player has blackjack, the stand method will be called from inside the hit method, which will call the finish dealer hand method
			return nil
		}
		game.dealerHit()
	}

	game.state = gameSessionState.StatePlayerTurn
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

	currentHandIndex := game.GetPlayer().GetCurrentHandIndex()

	_ = game.Hit()

	if currentHandIndex == game.GetPlayer().GetCurrentHandIndex() { //if player did not draw a blackjack or did not bust, stand automatically
		game.Stand()
	}

	return nil
}

func (game *game) Split() error {
	err := game.checkValidState(gameSessionState.StatePlayerTurn)

	if err != nil {
		return err
	}

	err = game.player.SplitHands()

	if err != nil {
		return err
	}
	game.Hit() //if split successful, the player has to be dealt another card, currently he has only one
	return nil
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
	if game.GetState() != gameSessionState.StatePlayerTurn {
		return fmt.Errorf(blackjackErrors.HitPlayerTurnError)
	}

	currentPlayerHand, err := game.getCurrentPlayerHand()
	if err != nil {
		return err
	}

	currentPlayerHand.AddCard(game.deck.DealCard())

	if currentPlayerHand.Score() >= 21 {
		game.Stand()
	}

	return nil
}

func (game *game) Stand() error {

	switch game.state {
	case gameSessionState.StatePlayerTurn:
		currentHandIndex := game.GetPlayer().GetCurrentHandIndex()
		if currentHandIndex == 0 {
			game.state = gameSessionState.StateDealerTurn
			game.FinishDealerHand()
		} else {
			game.GetPlayer().SetCurrentHandIndex(currentHandIndex - 1)
			game.Hit() // player splitted his hands, so he has only one card
		}
	case gameSessionState.StateDealerTurn:
		game.state = gameSessionState.StateHandOver
	default:
		return fmt.Errorf(blackjackErrors.NoActiveHands)
	}

	return nil
}

func (game *game) dealerHit() {
	game.GetDealer().GetDealerHandPointer().AddCard(game.deck.DealCard())
}

func (game *game) FinishDealerHand() error {
	err := game.checkValidState(gameSessionState.StateDealerTurn)
	if err != nil {
		return err
	}

	canHit := game.dealer.CanHit()

	for canHit { //while the dealer decides to hit, execute the hit method, crazy stuff i know
		game.dealerHit()
		canHit = game.dealer.CanHit()
	}

	game.Stand()

	return nil
}

func (game *game) EndHand() ([]outcome.BlackjackOutcome, outcome.Winnings, [][]outcome.MoneyOperation, error) { //TODO resolve insurance

	err := game.checkValidState(gameSessionState.StateHandOver)
	if err != nil {
		return nil, outcome.Winnings(0), [][]outcome.MoneyOperation{}, err
	}

	totalHands := game.GetPlayer().GetTotalHands()
	totalWinnings := outcome.Winnings(0)
	var handOutcomes []outcome.BlackjackOutcome
	var moneyOperations [][]outcome.MoneyOperation

	for i := uint8(0); i < totalHands; i++ {
		game.GetPlayer().SetCurrentHandIndex(i)

		handOutcome := outcome.ComputeOutcome(game.GetPlayer().GetCurrentHandCards(), game.GetDealer().GetDealerHand())
		game.GetPlayer().SetCurrentHandOutcome(handOutcome)

		handOutcomes = append(handOutcomes, handOutcome)

		winningsAmount, betBackAmount := outcome.ComputeWinningsForPlayer(handOutcome, game.GetPlayer().GetCurrentHandBet(), game.GetBlackjackPayout())

		game.GetPlayer().SetCurrentHandWinnings(winningsAmount + outcome.Winnings(betBackAmount))
		totalWinnings += winningsAmount + outcome.Winnings(betBackAmount)

		betAmount := game.GetPlayer().GetCurrentHandBet()
		moneyOperationsForCurrentHand, _ := outcome.ComputeMoneyOperationsForHand(betAmount, betBackAmount, winningsAmount)
		moneyOperations = append(moneyOperations, moneyOperationsForCurrentHand)
	}

	game.player.SetBalance(game.player.GetBalance() + int(totalWinnings))
	game.state = gameSessionState.StateBet

	min := 52 * game.numDecks / 3 //reshuffle after we consumed 2/3
	if len(game.deck) < min {
		game.ShuffleNewDeck()
	}
	game.player.ResetHands()
	game.dealer.ResetHands()

	return handOutcomes, totalWinnings, moneyOperations, nil
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
