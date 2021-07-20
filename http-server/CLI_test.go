package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/n4to4/learn-go-with-tests/http-server"
)

var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and start the game", func(t *testing.T) {
		in := strings.NewReader("7\n")
		stdout := &bytes.Buffer{}
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}
	})

	t.Run("finish game with 'Chris' as winner", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishedWith)
		}
	})

	t.Run("record 'Cleo' win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishedWith != "Cleo" {
			t.Errorf("expected finish called with 'Cleo' but got %q", game.FinishedWith)
		}
	})

	//t.Run("do not read beyond the first newline", func(t *testing.T) {
	//	in := failOnEndReader{
	//		t,
	//		strings.NewReader("1\nChris wins\nhello there"),
	//	}
	//})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		in := strings.NewReader("Pies\n")
		stdout := &bytes.Buffer{}
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

type GameSpy struct {
	StartedWith int
	StartCalled bool

	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

//type failOnEndReader struct {
//	t   *testing.T
//	rdr io.Reader
//}
//
//func (m failOnEndReader) Read(p []byte) (n int, err error) {
//	n, err = m.rdr.Read(p)
//
//	if n == 0 || err == io.EOF {
//		m.t.Fatal("Read to the end when you shouldn't have")
//	}
//
//	return n, err
//}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
