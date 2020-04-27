package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	// Decode the cipher!
	// Cipher only applies to [A-Za-z]
	const A_ascii byte = 65
	const Z_ascii byte = 90
	const a_ascii byte = 97
	const z_ascii byte = 122

	if b >= A_ascii && b <= Z_ascii {
		// Uppercase letters
		return (((b - A_ascii) + 13) % 26) + A_ascii

	} else if b >= a_ascii && b <= z_ascii {
		// Lowercase letters
		return (((b - a_ascii) + 13) % 26) + a_ascii

	} else {
		// This case will apply to spaces, punctuation, etc.
		return b // Just return argument
	}
}

func (z *rot13Reader) Read(p []byte) (n int, err error) {
	// Read bytes into p
	n, err = z.r.Read(p)

	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])

	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
