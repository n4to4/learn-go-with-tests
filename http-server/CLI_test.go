package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/n4to4/learn-go-with-tests/http-server"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		stdout := &bytes.Buffer{}

		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(&poker.SpyBlindAlerter{}, playerStore)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if len(playerStore.WinCalls) != 1 {
			t.Fatal("expected a win call but didn't get any")
		}

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		stdout := &bytes.Buffer{}

		playerStore := &poker.StubPlayerStore{}
		dummySpyAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(dummySpyAlerter, playerStore)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		stdout := &bytes.Buffer{}

		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, &poker.StubPlayerStore{})

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				poker.AssertScheduledAlert(t, got, want)
			})
		}
	})

	var dummyPlayerStore = &poker.StubPlayerStore{}

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")

		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				poker.AssertScheduledAlert(t, got, want)
			})
		}
	})
}
