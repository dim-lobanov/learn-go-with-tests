package arrays_and_slices

func Sum(arr []int) int {
	sum := 0

	for _, elem := range arr { // blank identifier instead of index
		sum += elem
	}

	return sum
}

// There's a new way to create a slice. `make` allows you to create a slice 
// with a starting capacity of the len of the numbersToSum we need to work through.
func SumAll(slices ...[]int) []int {
	var result []int // empty slice with 0 capacity, result[0] will panic with Out of Range runtime error
	// Note: `make([]int, 10)`` can create slice with non zero initial capacity
	// sl := make([]int, 10)
	// println(sl[9]) // good
	
	for _, slice := range slices {
		result = append(result, Sum(slice))
	}
	
	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

// If we want to pass array as a parameter, we have to specify fixed size
func SumArray(arr [5]int) (sum int) {
	for _, elem := range arr { // blank identifier instead of index
		sum += elem
	}
	return
}