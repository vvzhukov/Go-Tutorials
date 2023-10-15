/*
Packages, variables, and functions.
Learn the basic components of any Go program.
*/

package main

import ( //factored import statement
	"fmt"
	"math"
)

var ( //factored variables declaration
	// basic types
	stringvar string = "string"
	intvar    int    = 123       // int8  int16  int32  int64
	MaxInt    uint64 = 1<<64 - 1 // uint uint8 uint16 uint32 uint64 uintptr
	bytevar   byte               // alias for uint8
	runetype  rune               // alias for int32
	// represents a Unicode code point
	floattype   float64    // float32
	complextype complex128 // complex64
	ToBe        bool       = false
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
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

// non-initialized variables are initialized with their 'zero' values:
// 0 for int, false for bool, "" for string

const pi = 3.1415926 // somehow in tutorial all const are capitalized
const truth = true

func main() {
	// variables at function level
	var g int

	// implicit var declaration
	k := 2 // only inside functions
	// type inference (variable's type is inferred from the value on the right hand side)
	fmt.Printf("Type: %T Value: %v\n", k, k)

	fmt.Println(d, e, f, g, d2, e2, f2)
	fmt.Println(k)

	a, b := swap("Hello", "VZ")
	fmt.Println(b, a)

	fmt.Println("Custom function call: ", add(42, 3))
	fmt.Println("Exported names should be capitalized: ", math.Pi)

	fmt.Println(split(17))

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

	// type conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
