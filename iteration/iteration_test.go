package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertRepeatN := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeatN(t, repeated, expected)
	})

	t.Run("repeat 11 times", func(t *testing.T) {
		repeated := Repeat("b", 11)
		expected := "bbbbbbbbbbb"
		assertRepeatN(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 4)
	fmt.Println(repeated)
	// Output: aaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
