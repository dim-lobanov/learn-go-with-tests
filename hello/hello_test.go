package hello

import "testing"

// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T

func TestHello(t *testing.T) {
	// With t.Run we can run subtest in different goroutines
	t.Run("Hello World Default (subtest name)", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello World with Name", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Иван", "Russian")
		want := "Привет, Иван"
		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy,
// so you can call helper functions from a test, or a benchmark
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	t.Helper()
	if got != want {
		// %q (for strings) - a double-quoted string safely escaped with Go syntax
		// For tests %q is very useful as it wraps your values in double quotes.
		t.Errorf("Got %q, want %q\n", got, want) // Got "Goodbye, World", want "Hello, World"
		// t.Errorf("Got %v, want %v", got, want)   // Got Goodbye, World, want Hello, World
	}
}
