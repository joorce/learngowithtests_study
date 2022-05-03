package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(got, want)
	})

	t.Run("saying 'hello, world' when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(got, want)

	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(got, want)

	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"

		assertCorrectMessage(got, want)

	})

}
