package poker_test

import (
	. "server/poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Chris")
	})

    t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Cleo")
    })
}
