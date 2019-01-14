package outcome

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

func ComputeMoneyOperations(w Winnings, bet int) []MoneyOperation { //TODO

	return []MoneyOperation{{betBack, bet}}
}
