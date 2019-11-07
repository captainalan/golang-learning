package main

import "fmt"

const ARR_SIZE int = 20

// Pseudo-code from here: https://en.wikipedia.org/wiki/Insertion_sort
func insertion_sort(unsorted_list [ARR_SIZE]int) [ARR_SIZE]int {
	fmt.Println("Length of your slice", len(unsorted_list))
	var muh_list = [ARR_SIZE]int{} // Same size as unsorted list
	muh_list = unsorted_list       // Copy elements of unsorted_list argument

	// muh_list is modified in place, sorted one element at a time
	// with each iteration.
	for i := 1; i < len(muh_list); i++ {
		for j := i; j > 0 && muh_list[j-1] > muh_list[j]; j -= 1 {
			// Swapping
			temp := muh_list[j]
			muh_list[j] = muh_list[j-1]
			muh_list[j-1] = temp
		}
	}
	return muh_list
}

func main() {
	var unsorted_list = [ARR_SIZE]int{201, 39, 54, 461, 93, 377, 413, 176,
		164, 317, 109, 113, 36, 235, 17, 60, 53, 474, 257, 48}

	fmt.Println("Unsorted List", unsorted_list)
	fmt.Println("Sorted List", insertion_sort(unsorted_list))
}
