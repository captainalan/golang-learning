package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// "Exercise: Slices" from "A Tour of Go"
	// https://tour.golang.org/moretypes/18
	image := make([][]uint8, dy)
	for i, _ := range image {
		image[i] = make([]uint8, dx)

		// Apply some function to make a more interesting image
		for j, _ := range image[i] {
			image[i][j] = uint8((i + j) / 2)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
	// To view image, I piped the output of this program to some file (e.g. test.txt)
	// Then, I edited the IMAGE: prefix out of the beginning in an editor.
	// Finally, I used `base64 -d test.txt > test.png` to make a viewable image.
	// I'm patting myself on the back because I didn't StackOverflow or anything for halp! :)
}
