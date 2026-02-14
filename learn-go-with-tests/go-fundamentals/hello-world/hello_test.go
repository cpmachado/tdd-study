package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		// given
		name := "cpmachado"
		want := "hello, " + name

		// when
		got := Hello(name)

		// then
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'hello, world' when an empty string is supplied", func(t *testing.T) {
		// given
		name := ""
		want := "hello, world"

		// when
		got := Hello(name)

		// then
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
