package iteration

import "testing"

func TestRepeat(t *testing.T) {
	// given
	character := "a"
	n := 5
	expected := "aaaaa"

	// when
	repeated := Repeat(character, n)

	// then
	assertCorrectRepeat(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	character := "a"
	n := 5

	for b.Loop() {
		Repeat(character, n)
	}
}

func assertCorrectRepeat(t testing.TB, repeated, expected string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
