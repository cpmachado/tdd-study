package main

import "testing"

func TestHello(t *testing.T) {
	// given
	name := "cpmachado"
	want := "hello, " + name

	// when
	got := Hello(name)

	// then
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
