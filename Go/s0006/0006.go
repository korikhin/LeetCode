package main

import (
	"fmt"
	"strings"
)

func Convert(s string, numRows int) string {
	n := len(s)
	if s == "" || numRows == 1 || numRows >= n {
		return s
	}

	d := make([]strings.Builder, numRows)
	for _, b := range d {
		b.Grow(n/numRows + 1)
	}

	j, delta := 0, -1
	for _, r := range s {
		if j == 0 || j == numRows-1 {
			delta *= -1
		}
		d[j].WriteRune(r)
		j += delta
	}

	var z strings.Builder
	z.Grow(n)
	for _, b := range d {
		z.WriteString(b.String())
	}
	return z.String()
}

func main() {
	s := "PAYPALISHIRING"
	n := 3

	fmt.Println(Convert(s, n))
}
