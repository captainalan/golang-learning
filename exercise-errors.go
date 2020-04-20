package main

import (
	"fmt"
)

// Solution to https://tour.golang.org/methods/19

// Used in square root calculation
const EPSILON = 1e-5 // Set arbitrarily for now

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// hmm...
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Absolute(x float64) float64 {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func Sqrt(x float64) (float64, error) {
	// Algorithm as described in https://tour.golang.org/flowcontrol/8

	// Start with some guess $z$
	var z float64 = 1

	if x < 0 {
		// Here a 'type conversion' is done on x; ErrNegativeSqrt implements
		// error (built-in)
		return 0, ErrNegativeSqrt(x) // TODO what first value should I return here?
	}
	for {

		if Absolute(x-z*z) <= EPSILON {
			// If guess is good enough, return
			return z, nil
		} else {
			z -= (z*z - x) / (2 * z)
		}
	}

}

func main() {
	fmt.Println("Can we take the square root of 2?")
	fmt.Println(Sqrt(2))
	fmt.Println("Can we take the square root of -2?")
	fmt.Println(Sqrt(-2))
}
