// The caesar cipher.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func encode(s *bufio.Scanner, shift string) string {
	ns := len(shift)
	shift = strings.ToLower(shift)
	out := ""
	var ref rune
	var schar byte
	i := 0
	for _, v := range s.Text() {
		if v >= 'A' && v <= 'z' {
			if v > 'Z' {
				ref = 'a'
			} else {
				ref = 'A'
			}
			schar = shift[i] - 'a'
			v = ref + (v-ref+rune(schar))%26
			i = (i + 1) % ns
		}
		out += string(v)
	}

	return out
}

func main() {
	var shift string
	flag.StringVar(&shift, "shift", "cipher", "The shift")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	fmt.Println(encode(s, shift))

}
