package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	// a helper function that takes TB(which t and b fits) and compare the got and want value
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	// Check when language is English
	t.Run("saying hello to people in English", func(t *testing.T) {
		got := Hello("Edward", "English")
		want := "Hello, Edward"
		assertCorrectMessage(t, got, want)

	})

	// Check when language is Traditional Chinese
	t.Run("saying hello to people in Chinese", func(t *testing.T) {
		got := Hello("盈竹", "Traditional Chinese")
		want := "你好, 盈竹"

		assertCorrectMessage(t, got, want)
	})

	// Check when language is French
	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Edward", "French")
		want := "Bonjour, Edward"
		assertCorrectMessage(t, got, want)
	})

	// Check when language is Japanese
	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("盈竹", "Japanese")
		want := "こんにちは, 盈竹さん"

		assertCorrectMessage(t, got, want)
	})

	// Check when input name is empty
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	// Check when input language is not defined, empty or gibberish
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		// not defined
		got := Hello("Edward", "Thai")
		want := "Hello, Edward"
		assertCorrectMessage(t, got, want)
		// empty
		got = Hello("Edward", "")
		assertCorrectMessage(t, got, want)
		//  gibberish
		got = Hello("Edward", "anyGibberish")
		assertCorrectMessage(t, got, want)
	})

}

func ExampleHello() {
	greetings := Hello("Edward", "")
	fmt.Println(greetings)
	// Output: Hello, Edward

}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("", "")
	}
}
