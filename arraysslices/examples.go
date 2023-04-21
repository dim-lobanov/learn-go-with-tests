package arraysslices

import "fmt"

// Changing the slice affects the original array; but a "copy" of the slice will not affect the original array
func UnderlyingArrayChangesExample() {
	x := [3]string{"Лайка", "Белка", "Стрелка"}

	y := x[:] // slice "y" points to the underlying array "x"

	z := make([]string, len(x))
	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "Belka" // the value at index 1 is now "Belka" for both "y" and "x"

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}

// Why it's a good idea to make a copy of a slice after slicing a very large slice
func CopyingSliceOfLargeArrayExample() {
	a := make([]int, 1e6) // slice "a" with len = 1 million
	b := a[:2]            // even though "b" len = 2, it points to the same the underlying array "a" points to

	c := make([]int, len(b)) // create a copy of the slice so "a" can be garbage collected
	copy(c, b)
	fmt.Println(c)
}
