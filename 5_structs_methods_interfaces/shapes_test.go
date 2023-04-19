package structs_methods_interfaces

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	t.Run("Rectagles Perimeter", func(t *testing.T) {
		got := Perimeter(Rectangle{10.0, 10.0})
		expected := 40.0

		if got != expected {
			t.Errorf("got %.2f, want %.2f", got, expected)
		}
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()
		if got != expected {
			t.Errorf("got %g want %g", got, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		checkArea(t, Rectangle{12, 6}, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		checkArea(t, Circle{10}, 314.1592653589793)
	})
}

// https://github.com/golang/go/wiki/TableDrivenTests
func TestAreaWithTable(t *testing.T) {
	// "anonymous struct"
	areaTests := []struct {
		shape    Shape
		expectedArea float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{shape: Triangle{base: 12, height: 6}, expectedArea: 36.0}, // We can explicitly specify names for structs fields
		// Note: we cannot mix explicit and implicit declaration inside struct
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.expectedArea {
			t.Errorf("got %g, expected area %g", got, tt.expectedArea)
		}
	}
}

// In Go interface resolution is implicit.
// If the type you pass in matches what the interface is asking for, it will compile.
