package main

import "strings"

func Convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}

	b := make([]strings.Builder, numRows)
	c := 2*(n-1)/(2*numRows-2) + 1
	for i := range b {
		b[i].Grow(c)
	}

	i, delta := 0, -1
	for _, r := range s {
		if i == 0 || i == numRows-1 {
			delta *= -1
		}
		b[i].WriteRune(r)
		i += delta
	}

	var z strings.Builder
	z.Grow(n)
	for _, bb := range b {
		z.WriteString(bb.String())
	}
	return z.String()
}
