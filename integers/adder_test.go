package integers

// We can declare aliases for imports
import test "testing"

func TestAdder(t *test.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d, but got %d", expected, sum)
	}
}