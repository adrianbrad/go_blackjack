package blackjackErrors

const (
	NoActiveHands             = "currently there is no players turn"
	NoBetsPlaced              = "no bets placed"
	BetAlreadyPlaced          = "bet already placed"
	BetHigherThanBalance      = "bet given higher then balance"
	BetHasToBeGreaterThanZero = "bet has to be greater than 0"
	DealerFirstCardError      = "dealer first has to be an Ace"
	SplitCardsNumberError     = "you can split only with two cards in hand"
	SplitCardsValueError      = "you can split only with two cards of the same value"
)
