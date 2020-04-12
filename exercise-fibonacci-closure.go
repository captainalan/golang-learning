package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// value1 initialized at 0 indicates start of sequence
	var value1, value2 int = 0, 0

	f := func() int {
		var to_return int
		if value1 == 0 && value2 == 0{ // handles the "base case"
			to_return = 1;
			value1 = 1;
		} else {
			to_return = value1 + value2
		}
		value1 = value2
		value2 = to_return
		return to_return;
	}
	return f
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
