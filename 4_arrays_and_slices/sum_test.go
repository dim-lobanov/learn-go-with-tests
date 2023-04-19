package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 9})
	expected := []int{6, 9}

	// You could write a function to iterate over each got and want slice and check their values
	// Or use reflect.DeepEqual
	// It's important to note that reflect.DeepEqual is not "type safe" -
	// the code will compile even if you did something stupid
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %d want %d given", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	// local fucntion for reduced scope
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

// Go's built-in testing toolkit features a coverage tool
// `go test -cover`
