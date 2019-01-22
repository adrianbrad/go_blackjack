package blackjackWinner

type BlackjackWinner uint8

const (
	Player BlackjackWinner = iota
	Dealer
	Draw
)

var winners = [...]string{"Player Won", "Dealer Won", "Draw"}

func (w BlackjackWinner) String() string {
	return winners[w] //[...] - specifies the length is equal to the number of elements in the array literal
}
