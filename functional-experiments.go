package main

import "fmt"

/* 
TODO
- Implement map, filter, and (right) reduce
- Rather than just using ints, use interfaces or something to accept more
  generic types with these functions
*/

// Returns a function like
func add_function_maker(n int) func(int) int{
	return func(m int) int{
		return m + n;
	}
}

/*
func map(){

}
*/

func main() {
	fmt.Println("hello world")

	incrementer := add_function_maker(1)
	fmt.Println("Should be 3:", incrementer(2))
}

