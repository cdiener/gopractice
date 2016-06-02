// The caesar cipher.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func encode(s *bufio.Scanner, shift int) string {
	out := ""
	var ref rune
	for _, v := range s.Text() {
		if v >= 'A' && v <= 'z' {
			if v > 'Z' {
				ref = 'a'
			} else {
				ref = 'A'
			}
			v = ref + (v-ref+rune(shift))%26
		}
		out += string(v)
	}

	return out
}

func main() {
	var shift int
	flag.IntVar(&shift, "shift", 13, "The shift")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	fmt.Println(encode(s, shift))

}
