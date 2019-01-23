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
	NoMoneyForDoubleDown      = "not enough money to double down"
	NoMoneyForSplitting       = "not enough money to split"
	InvalidCardsForDoubleDown = "current hand has to have 2 cards"
	BetAlreadyDoubled         = "bet already doubled"
	InvalidInsuranceHand      = "invalid hand for insurance"
	InvalidBalance            = "invalid balance for operation"
	InvalidSetWinningsHand    = "invalid hand for settings winnings"
	InvalidSetOutcomeHand     = "invalid hand for settings outcome"
	HitPlayerTurnError        = "cannot hit if it is not player's turn"
)
