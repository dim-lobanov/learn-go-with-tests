package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} // Buffer implements io.Writer interface
	Greet(&buffer, "Chris")

	got := buffer.String() // Buffer implements Stringer interface { String() string }
	expected := "Hello, Chris"

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}

// type Writer interface {
// 	  Write(p []byte) (n int, err error)
// }
