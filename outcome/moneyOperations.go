package outcome

import "fmt"

type MoneyOperationType uint8

type MoneyOperation struct {
	OperationType MoneyOperationType
	amount        int
}

const (
	bet MoneyOperationType = iota
	betBack
	win
)

var operations = [...]string{"Player bet amount", "Bet back amount", "Winnings amount"}

func (m MoneyOperationType) String() string {
	return operations[m] //[...] - specifies the length is equal to the number of elements in the array literal
}

//ComputeMoneyOperationsForHand used to describe the money flow
func ComputeMoneyOperationsForHand(betAmount int, betBackAmount BetBack, winningsAmount Winnings) ([]MoneyOperation, error) {
	var moneyOperations []MoneyOperation

	if betAmount == 0 {
		return moneyOperations, fmt.Errorf("Bet cannot be 0")
	}
	moneyOperations = append(moneyOperations, MoneyOperation{bet, betAmount})

	if betBackAmount != 0 {
		moneyOperations = append(moneyOperations, MoneyOperation{betBack, int(betBackAmount)})
	}

	if winningsAmount != 0 {
		moneyOperations = append(moneyOperations, MoneyOperation{win, int(winningsAmount)})
	}

	return moneyOperations, nil
}
