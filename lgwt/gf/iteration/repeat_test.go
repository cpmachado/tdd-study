package iteration

import "testing"

func TestRepeat(t *testing.T) {
	// given
	character := "a"
	expected := "aaaaa"

	// when
	repeated := Repeat(character)

	// then
	assertCorrectRepeat(t, repeated, expected)
}

func assertCorrectRepeat(t testing.TB, repeated, expected string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
