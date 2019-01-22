package game

import (
	"blackjack/dealer"
	"blackjack/gameSessionState"
	"blackjack/hand"
	"blackjack/outcome"
	"blackjack/player"
	"reflect"
	"testing"

	deck "github.com/adrianbrad/go-deck-of-cards"
)

func TestNew(t *testing.T) {
	type args struct {
		numDecks        int
		blackjackPayout float64
		player          player.Player
		dealer          dealer.Dealer
		deck            deck.Deck
	}
	tests := []struct {
		name string
		args args
		want game
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.numDecks, tt.args.blackjackPayout, tt.args.player, tt.args.dealer, tt.args.deck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_Bet(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	type args struct {
		bet int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.Bet(tt.args.bet); (err != nil) != tt.wantErr {
				t.Errorf("game.Bet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_game_PlaceInsurance(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.PlaceInsurance(); (err != nil) != tt.wantErr {
				t.Errorf("game.PlaceInsurance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_game_DoubleDown(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.DoubleDown(); (err != nil) != tt.wantErr {
				t.Errorf("game.DoubleDown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_game_Split(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.Split(); (err != nil) != tt.wantErr {
				t.Errorf("game.Split() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_game_getCurrentPlayerHand(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		want    *hand.Hand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			got, err := game.getCurrentPlayerHand()
			if (err != nil) != tt.wantErr {
				t.Errorf("game.getCurrentPlayerHand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("game.getCurrentPlayerHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_ShuffleNewDeck(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			game.ShuffleNewDeck()
		})
	}
}

func Test_game_Hit(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.Hit(); (err != nil) != tt.wantErr {
				t.Errorf("game.Hit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_game_Stand(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.Stand(); (err != nil) != tt.wantErr {
				t.Errorf("game.Stand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_game_DealStartingHands(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.DealStartingHands(); (err != nil) != tt.wantErr {
				t.Errorf("game.DealStartingHands() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_game_dealerHit(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			game.dealerHit()
		})
	}
}

func Test_game_FinishDealerHand(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.FinishDealerHand(); (err != nil) != tt.wantErr {
				t.Errorf("game.FinishDealerHand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_game_EndHand(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []outcome.BlackjackOutcome
		want1   outcome.Winnings
		want2   []outcome.MoneyOperation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			got, got1, got2, err := game.EndHand()
			if (err != nil) != tt.wantErr {
				t.Errorf("game.EndHand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("game.EndHand() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("game.EndHand() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("game.EndHand() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_game_GetDealer(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name   string
		fields fields
		want   dealer.Dealer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if got := game.GetDealer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("game.GetDealer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_GetDeck(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name   string
		fields fields
		want   deck.Deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if got := game.GetDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("game.GetDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_GetPlayer(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name   string
		fields fields
		want   player.Player
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if got := game.GetPlayer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("game.GetPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_GetState(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name   string
		fields fields
		want   gameSessionState.GameSessionState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if got := game.GetState(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("game.GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_GetBlackjackPayout(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if got := game.GetBlackjackPayout(); got != tt.want {
				t.Errorf("game.GetBlackjackPayout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_checkValidState(t *testing.T) {
	type fields struct {
		numDecks        int
		blackjackPayout float64
		deck            deck.Deck
		initialDeck     deck.Deck
		state           gameSessionState.GameSessionState
		player          player.Player
		currentPlayer   uint8
		totalPlayers    uint8
		dealer          dealer.Dealer
	}
	type args struct {
		givenState gameSessionState.GameSessionState
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := game{
				numDecks:        tt.fields.numDecks,
				blackjackPayout: tt.fields.blackjackPayout,
				deck:            tt.fields.deck,
				initialDeck:     tt.fields.initialDeck,
				state:           tt.fields.state,
				player:          tt.fields.player,
				currentPlayer:   tt.fields.currentPlayer,
				totalPlayers:    tt.fields.totalPlayers,
				dealer:          tt.fields.dealer,
			}
			if err := game.checkValidState(tt.args.givenState); (err != nil) != tt.wantErr {
				t.Errorf("game.checkValidState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
