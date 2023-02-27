package poker_test

import (
	"bytes"
	"server/poker"
	. "server/poker"
	"strings"
	"testing"
)

var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		game := &GameSpy{}

		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("wanted Finish called with Chris but got %v", game.FinishedWith)
		}
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		game := &GameSpy{}

		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Cleo" {
			t.Errorf("wanted Finish called with Cleo but got %v", game.FinishedWith)
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("it prints an error when a non-numeric value is entered and does not start the game", func(t *testing.T) {
        stdout := &bytes.Buffer{}
        in := strings.NewReader("Pies\n")
        game := &GameSpy{}

        cli := poker.NewCLI(in, stdout, game)
        cli.PlayPoker()

        if game.StartCalled {
            t.Errorf("game should not have started")
        }

        gotPrompt := stdout.String()
        wantPrompt := poker.PlayerPrompt + "Bad value received for number of players, please try again with a number"

        if gotPrompt != wantPrompt {
            t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
        }
	})
}
