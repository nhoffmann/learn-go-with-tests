package poker_test

import (
	"strings"
	"testing"

	poker "github.com/nhoffmann/learn-go-with-tests/command-line"
)

func TestCLI(t *testing.T) {
	t.Run("Record Niels win from user input", func(t *testing.T) {
		in := strings.NewReader("Niels wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Niels")
	})

	t.Run("Record Chloe win from user input", func(t *testing.T) {
		in := strings.NewReader("Chloe wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chloe")
	})
}
