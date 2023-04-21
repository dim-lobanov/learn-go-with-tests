package structsinterfaces

import "math"

// struct is just a named collection of fields where you can store data
type Rectangle struct {
	width  float64
	height float64
}

// A method is a function with a receiver
// (r Rectangle) - receiver
// func (receiverName ReceiverType) MethodName(args)
// receiverName inside method works like 'this' in other languages
//
// It is a convention in Go to have the receiver variable be the first letter of the type.
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}

// There is no function overloading in Go

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}
