package client

import (
	"blackjack/game"
	"blackjack/gameSessionState"
	"fmt"
	"strconv"
	"strings"
)

//Client interface holds the methods for a client api 
type Client interface {
	DisplayOptions()
}

type client struct {
	givenInput string
	game       game.Game
}

func NewCLIClient(game game.Game) *client {
	g := &client{
		game: game,
	}
	separator()
	fmt.Println("New game started.")
	fmt.Println("Player balance", g.game.GetPlayer().GetBalance())
	separator()

	return g
}

func (c client) DisplayOptions() {
	fmt.Println()
	separator()
	fmt.Println("Choose an option:")
	fmt.Println("(d)eal; (b)et*amount*; (h)it; (s)tand; s(p)lit; (i)nsurance; d(o)ubledown; (a)dd funds*amount*; add pla(y)er")
	separator()
	fmt.Println()
	c.askForInput()
}

func separator() {
	fmt.Println(strings.Repeat("*", 10))
}

func (c *client) bet(betInput string) error {
	if len(betInput) < 1 {
		return fmt.Errorf("invalid bet amount")
	}
	amount, err := strconv.Atoi(betInput[1:])
	if err != nil {
		return fmt.Errorf("invalid bet format")
	}

	err = c.game.Bet(amount)
	if err == nil {
		fmt.Println("Successfully bet: ", c.game.GetPlayer().GetCurrentHandBet())
		fmt.Println("Balance: ", c.game.GetPlayer().GetBalance())
	}
	return err
}

func (c *client) deal() error {
	err := c.game.DealStartingHands()
	if err != nil {
		return err
	}
	c.displayGameInfo()

	return nil
}

func (c client) displayGameInfo() {
	separator()
	fmt.Println("Bet: ", c.game.GetPlayer().GetCurrentHandBet())
	fmt.Println("Balance: ", c.game.GetPlayer().GetBalance())
	c.displayPlayerHandsInfo()
	c.displayDealerHandInfo()
	separator()
}

func (c client) displayPlayerHandsInfo() {
	playerHands := c.game.GetPlayer().GetHands()
	fmt.Println("Player hands")
	for idx,hand := range playerHands {
		fmt.Println(idx + 1, " ", hand, " score: ", hand.Score())
	}

}
func (c client) displayDealerHandInfo() {
	fmt.Println("Dealer hand")
	if c.game.GetState() != gameSessionState.StateHandOver {
		dealerFirstCard,_ := c.game.GetDealer().GetDealerFirstCard()
		fmt.Println(dealerFirstCard, "Score: ", dealerFirstCard.Score())
	} else {
		dealerHand := c.game.GetDealer().GetDealerHand()
		fmt.Println(dealerHand, "Score: ", dealerHand.Score())
	}
}

func (c *client) hit() error {
	err := c.game.Hit()
	if err != nil {
		return err
	}

	c.displayGameInfo()
	c.finishGameIfPossible()
	return nil
}

func (c *client) stand() error {
	err := c.game.Stand()
	if err != nil {
		return err
	}
	c.displayGameInfo()
	c.finishGameIfPossible()
	return nil
}

func (c *client) split() error {
	err := c.game.Split()
	if err != nil {
		return err
	}
	c.displayGameInfo()
	return nil
}

func (c *client) insurance() error {
	err := c.game.PlaceInsurance()
	if err != nil {
		return err
	}
	c.displayGameInfo()
	return nil
}

func (c *client) doubledown() error {
	err := c.game.DoubleDown()
	if err != nil {
		return err
	}
	c.displayGameInfo()
	c.finishGameIfPossible()
	return nil
}

func (c *client) finishGameIfPossible() {
	if c.game.GetState() == gameSessionState.StateHandOver {
		outcome, winnings, moneyOp, _ := c.game.EndHand()
		fmt.Println("Outcome: ", outcome)
		fmt.Println("Winnings: ", winnings)
		fmt.Println("Money operations: ", moneyOp)
		fmt.Println("Balance: ", c.game.GetPlayer().GetBalance())
	}
}

func (c *client) askForInput() {
	var input string
	var err error
	_, _ = fmt.Scanf("%s\n", &input)

	if len(input) > 0 {
		c.givenInput = input[0:1]
		switch c.givenInput {
		case "d":
			err = c.deal()
		case "b":
			err = c.bet(input)
		case "h":
			err = c.hit()
		case "s":
			err = c.stand()
		case "p":
			err = c.split()
		case "i":
			err = c.insurance()
		case "o":
			err = c.doubledown()
		case "a":
			//c.game.
		case "y":
			//c.game.
		default:
			fmt.Printf("Current given input: %s, is invalid", input)
		}
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	c.DisplayOptions()
}
