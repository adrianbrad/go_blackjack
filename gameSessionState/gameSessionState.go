package gameSessionState

type GameSessionState uint8

const (
	StateBet GameSessionState = iota
	StatePlayerTurn
	StateDealerTurn
	StateHandOver
)

var suits = [...]string{"StateBet", "StatePlayerTurn", "StateDealerTurn", "StateHandOver"}

func (s GameSessionState) String() string {
	return suits[s]
}
