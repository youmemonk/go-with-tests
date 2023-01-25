package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris", "English")
	want := "Hello, Chris"
	assertCorrectMessage(t, got, want)

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func (t *testing.T)  {
		got := Hello("Loid", "Spanish")
		want := "Hola, Loid"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func (t *testing.T){
		got := Hello("Frenchie", "French")
		want := "Bonjour, Frenchie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
