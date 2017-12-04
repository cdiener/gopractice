package main

import (
	"flag"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	num := strings.Join(flag.Args(), "")
	digits := strings.Split(num, "")
	n := len(digits)
	sum := 0

	for i := range digits {
		if digits[i] == digits[(i+n/2)%n] {
			d, error := strconv.ParseInt(digits[i], 0, 64)
			if error != nil {
				panic("Could not parse digit!")
			}
			sum += int(d)
		}
	}

	println(sum)

}
