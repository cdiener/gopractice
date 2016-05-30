// The caesar cipher.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var shift int

func main() {
	flag.IntVar(&shift, "shift", 13, "The shift")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	s.Scan()

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
		fmt.Printf("%c", v)
	}
	fmt.Println()

}
