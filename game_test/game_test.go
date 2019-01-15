package game_test

import (
	"blackjack/dealer"
	"blackjack/game"
	"blackjack/gameSessionState"
	"blackjack/player"
	"testing"
)

func TestNewGame(t *testing.T) {
	var g game.Game

	g = game.New(5, 1.5, player.New(30), dealer.NewDefaultDealer())

	equals(t, g.GetState(), gameSessionState.StateBet)
	equals(t, g.GetBlackjackPayout(), 1.5)
	equals(t, len(g.GetDeck()), 52*5)

	equals(t, g.GetPlayer().GetCurrentHandBet(), 0)
	equals(t, len(g.GetPlayer().GetCurrentHandCards()), 0)
	equals(t, g.GetPlayer().GetBalance(), 30)
	equals(t, g.GetPlayer().GetTotalHands(), uint8(1))
	equals(t, g.GetPlayer().GetCurrentHandIndex(), uint8(0))

	_, err := g.GetDealer().GetDealerFirstCard()
	equals(t, err.Error(), "dealer has no cards")
	equals(t, len(g.GetDealer().GetDealerHand()), 0)
}

func TestBet(t *testing.T) {

}
