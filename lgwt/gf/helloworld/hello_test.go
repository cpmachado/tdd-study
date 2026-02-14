package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		// given
		name := "cpmachado"
		language := "English"
		want := "Hello, " + name

		// when
		got := Hello(name, language)

		// then
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		// given
		name := ""
		language := "English"
		want := "Hello, World"

		// when
		got := Hello(name, language)

		// then
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		// given
		name := "Elodie"
		language := "Spanish"
		want := "Hola, Elodie"

		// when
		got := Hello(name, language)

		// then
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		// given
		name := "Elodie"
		language := "French"
		want := "Bonjour, Elodie"

		// when
		got := Hello(name, language)

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
