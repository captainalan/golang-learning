package main
import (
	"fmt"
)
/* Following the algorithm given under 'Procedure' here
   https://en.wikipedia.org/wiki/Binary_search_algorithm
*/

func binary_search(sorted_list []int, target int) int {
	// Assumes sorted_list is sorted
	// target is the thing you are searching for

	// Left and Right bounds to search
	var L, R int = 0, len(sorted_list) -1 
	// Position of the middle element
	var m int 

	// Find golang idiomatic way to write this loop
	for ; ; {
		m = (L + R)/2
		if sorted_list[m] < target {
			L = m + 1
		} else if sorted_list[m] > target {
			R = m - 1
		} else {
			return m;
		}
		if L > R { // Condition to break loop
			return -1;
		} 
	}
}

// Implementation of binary sort in go 
func main() {
	sorted_list := []int{-999, 1, 2, 3, 4, 5}
	target      := 2
	fmt.Printf("Looking for %v in %v.\n", target, sorted_list)
	fmt.Printf("The result? Index is %v\n", binary_search(sorted_list, 2))
}
