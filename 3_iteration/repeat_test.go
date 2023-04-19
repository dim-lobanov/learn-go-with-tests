package iteration

import . "testing"

func TestRepeat(t *T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q, but got %q", expected, repeated)	
	}
}

// To run the benchmarks do `go test -bench=.` (or if you're in Windows Powershell `go test -bench="."`)
func BenchmarkRepeat(b *B) {
	for i := 0; i < b.N; i++ { // framework determines "good" amount of iterations for benchmark - b.N
		Repeat("a", 5)
	}
}

// BenchmarkRepeat-20
// 17038290                66.20 ns/op           16 B/op          4 allocs/op
// Benchmark was run 17038290 times, 66.20 nanoseconds to run on average
//
// NOTE by default Benchmarks are run sequentially.