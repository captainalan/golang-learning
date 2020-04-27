package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (z *rot13Reader) Read(p []byte) (n int, err error) {
	// Read bytes into p
	n, err = z.r.Read(p) // TODO figure out how to get this to NOT print

	// Decode the cipher!
	// Cipher only applies to [A-Za-z]
	const A_ascii byte = 65
	const Z_ascii byte = 90
	const a_ascii byte = 97
	const z_ascii byte = 122

	for i := 0; i < n; i++ {
		var current byte = p[i] // Current letter (type int)
		if current >= A_ascii && current <= Z_ascii {
			// Uppercase letters
			rot13 := (((p[i] - A_ascii) + 13) % 26) + A_ascii
			fmt.Printf("%c", rot13)

		} else if current >= a_ascii && current <= z_ascii {
			// Lowercase letters
			rot13 := (((p[i] - a_ascii) + 13) % 26) + a_ascii
			fmt.Printf("%c", rot13)

		} else {
			// This case will apply to spaces, punctuation, etc.
			fmt.Printf("%c", p[i])
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
