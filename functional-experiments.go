package main

import "fmt"

/* 
TODO
- Implement (right) reduce
- Rather than just using ints, use interfaces or something to accept more
  generic types with these functions
*/
func add_function_maker(n int) func(int) int{
	return func(m int) int{
		return m + n;
	}
}

func my_map(my_func func(int) int, my_list []int) []int {
	// Create new slice of same size as input slice
	to_return := make([]int, len(my_list))
	for i := 0; i < len(my_list); i++ {
		to_return[i] = my_func(my_list[i])
	}
	return to_return;
}

func my_filter(my_predicate func(int) bool, my_list []int) []int {
	// Create new slice of same size as input slice
	to_return := make([]int, len(my_list))

	var current_position int = 0
	for i := 0; i < len(my_list); i++ {
		if my_predicate(my_list[i]) {
			to_return[current_position] = my_list[i]
			current_position++
		} else {
			// Do nothing
		}
	}
	// Return new slice, cutting off the unused space
	return to_return[:current_position];
}

const ARR_SIZE = 10

func main() {
	fmt.Println("Some basic higher order functions")

	// Sequence of numbers 1-10
	// TODO find a better way to get sequences
	var my_arr = [ARR_SIZE]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original array", my_arr) 

	// Create a function that increments by 5
	plus_five := add_function_maker(5)
	fmt.Println("Created a function that adds five to ints...") 

	// Map it!
	five_added := my_map(plus_five, my_arr[:])
	fmt.Println("...and mapped this to a slice of our original array.", five_added)

	// Create a function that returns true if int argument is even 
	var evenp func(int) bool = func(num int) bool {
		return num % 2 == 0
	}
	fmt.Println("Created a function that returns true for even integers") 

	// Filter it!
	evens := my_filter(evenp, my_arr[:])
	fmt.Println("...and mapped this to a slice of our original array.", evens)
}

