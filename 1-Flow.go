/*
Learn how to control the flow of your code with conditionals, loops, switches and defers.
*/

package main

import ( // factored import statement
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 { // expression need not be surrounded by parentheses
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // short statement to execute before condition
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	} // here ends the scope for the 'v'
	return lim
}

func sqrt2(x float64, max_iter int, precision float64) float64 {

	if precision == 0.0 {
		precision = 0.000000001
	}

	z := 1.0
	prev := 0.0

	for i := 1; i < max_iter; i++ {
		z -= (z*z - x) / (2 * z) // Newton's method
		fmt.Println(z)

		if math.Abs(z-prev) < precision {
			fmt.Println("Caught!")
			return z
		}
		prev = z

	}
	return z
}

func main() {

	// ----------------------------------------------------
	sum := 0                  // implicit declaration
	for i := 0; i < 10; i++ { // the only loop in go lang
		sum += i
	}

	// 3 statements: init / condition / post
	// first and last are optional

	sum2 := 1

	for sum2 < 1000 { // that is similar to while, wow :)
		sum2 += sum2
	}

	/*
		// and here is 'while true'
		for {

		}
	*/

	//fmt.Println(sum, sum2)

	// ----------------------------------------------------

	//fmt.Println("sqrt output: ", sqrt(2), sqrt(-4))
	//fmt.Println("pow output: ", pow(3, 2, 10), pow(3, 3, 20))
	//fmt.Println(sqrt2(2, 1000, 0.001))

	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch evaluation from top to bottom, 'breaks' when succeed
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch with condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//  defers the execution of a function until the surrounding function returns
	defer fmt.Println("world")
	fmt.Println("hello")

	// defer calls are pushed onto stack

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
