// Test the caesar cipher
package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCaesar(t *testing.T) {
	cases := []struct{ in, want, shift string }{
		{"Hello, world!", "Jmass, nqzak!", "cipher"},
		{"Hello, world!", "Hello, world!", "a"},
		{"Hello, world!", "Hello, world!", "A"},
	}

	var s string
	var scan *bufio.Scanner
	for _, set := range cases {
		scan = bufio.NewScanner(strings.NewReader(set.in))
		scan.Scan()
		s = encode(scan, set.shift)
		if s != set.want {
			t.Errorf("'%s' did encode to '%s' and not to '%s'!", set.in, s, set.want)
		}
	}
}
