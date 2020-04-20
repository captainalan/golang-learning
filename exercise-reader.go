package main

import  "golang.org/x/tour/reader"
// Solution to Exercise: Readers
// https://tour.golang.org/methods/22

type MyReader struct{}

// Working from go source: https://golang.org/src/strings/reader.go?s=1042:1092#L28
func (m MyReader) Read(p []byte) (n int, err error) {
	// For length of byte assignment, set things to 'A'
	for i, _ := range p {
		p[i] = byte('A')
	}
	// Finall, return 'A' (rather than EOF, as we are returning 'A's forever)
	return int('A'), nil
}

func main() {
	reader.Validate(MyReader{})
}
