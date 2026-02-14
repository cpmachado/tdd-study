package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	// given
	character := "a"
	count := 5
	expected := "aaaaa"

	// when
	repeated := Repeat(character, count)

	// then
	assertCorrectRepeat(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	character := "a"
	count := 5

	for b.Loop() {
		Repeat(character, count)
	}
}

// [strings.Repeat] is identic to our function and runs in similar time
func BenchmarkStringsRepeat(b *testing.B) {
	character := "a"
	count := 5

	for b.Loop() {
		strings.Repeat(character, count)
	}
}

func ExampleRepeat() {
	character := "a"
	count := 5
	fmt.Println(Repeat(character, count))

	// Output: aaaaa
}

func assertCorrectRepeat(t testing.TB, repeated, expected string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
