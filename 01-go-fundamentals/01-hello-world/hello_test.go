package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("you", "")
		want := "hello, you"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("maria", "spanish")
		want := "hola, maria"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("celine", "french")
		want := "bonjour, celine"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in portuguese", func(t *testing.T) {
		got := Hello("gabriel", "portuguese")
		want := "olá, gabriel"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
