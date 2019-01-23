package hand_test

import (
	"blackjack/hand"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func equals(tb testing.TB, act, exp interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
func TestHandScore(t *testing.T) {
	//Two cards hands
	equals(t, hand.Hand{{1, 2}, {1, 3}}.Score(), 5)
	equals(t, hand.Hand{{1, 1}, {1, 1}}.Score(), 12)
	equals(t, hand.Hand{{1, 11}, {1, 12}}.Score(), 20)
	equals(t, hand.Hand{{1, 1}, {1, 12}}.Score(), 21)

	//Three cards hands
	equals(t, hand.Hand{{1, 2}, {1, 2}, {1, 2}}.Score(), 6)
	equals(t, hand.Hand{{1, 1}, {1, 1}, {1, 12}}.Score(), 12)
	equals(t, hand.Hand{{1, 1}, {1, 1}, {1, 1}}.Score(), 13)
	equals(t, hand.Hand{{1, 7}, {1, 8}, {1, 9}}.Score(), 24)
	equals(t, hand.Hand{{1, 11}, {1, 12}, {1, 13}}.Score(), 30)

	//Four cards hands
	equals(t, hand.Hand{{1, 1}, {1, 1}, {1, 1}, {1, 12}}.Score(), 13)
	equals(t, hand.Hand{{1, 1}, {1, 1}, {1, 1}, {1, 1}}.Score(), 14)
	equals(t, hand.Hand{{1, 5}, {1, 5}, {1, 5}, {1, 5}}.Score(), 20)
	equals(t, hand.Hand{{1, 7}, {1, 8}, {1, 9}, {1, 13}}.Score(), 34)
	equals(t, hand.Hand{{1, 10}, {1, 11}, {1, 12}, {1, 13}}.Score(), 40)
}

func TestBlackjackHand(t *testing.T) {
	//TRUE
	equals(t, hand.Hand{{1, 1}, {1, 10}}.Blackjack(), true)
	equals(t, hand.Hand{{1, 11}, {1, 1}}.Blackjack(), true)

	//FALSE
	equals(t, hand.Hand{{1, 2}, {1, 10}}.Blackjack(), false)
	equals(t, hand.Hand{{1, 1}, {1, 10}, {1, 10}}.Blackjack(), false)
	equals(t, hand.Hand{{1, 1}, {1, 1}}.Blackjack(), false)
}

func TestHandIndividualValues(t *testing.T) {
	equals(t, hand.Hand{{1, 1}, {1, 10}}.IndividualScore(), []int{11, 10})
	equals(t, hand.Hand{{1, 1}, {1, 12}, {1, 1}}.IndividualScore(), []int{11, 10, 1})
	equals(t, hand.Hand{{1, 11}, {1, 12}, {1, 13}}.IndividualScore(), []int{10, 10, 10})
}
