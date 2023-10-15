package main

import ( //factored import statement
	"fmt"
	"math"
)

var ( //factored variables declaration
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
)

func add(x, y int) int {
	return x + y
}

// functions with multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// functions with 'naked' return
// only use in short functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// variables at package level
var d, e, f bool = false, true, false

// we may omit var type when var is initialized
var d2, e2, f2 = false, "string type", false

func main() {
	// variables at function level
	var g int

	// implicit var declaration
	k := 2 // only inside functions

	fmt.Println(d, e, f, g, d2, e2, f2)
	fmt.Println(k)
	a, b := swap("Hello", "VZ")
	fmt.Println(b, a)
	fmt.Println("Custom function call: ", add(42, 3))
	fmt.Println("Exported names should be capitalized: ", math.Pi)
	fmt.Println(split(17))
}
