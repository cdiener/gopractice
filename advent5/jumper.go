package main

import (
	"flag"
	"strconv"
	"strings"
)

func toInt(input string) []int {
	fields := strings.Fields(input)
	n := len(fields)
	ints := make([]int, n)
	for i := range ints {
		ints[i], _ = strconv.Atoi(fields[i])
	}

	return ints
}

func main() {
	flag.Parse()
	offsets := toInt(flag.Args()[0])
	n := len(offsets)
	loc := 0
	steps := 0

	for {
		if loc < 0 || loc >= n {
			break
		}
		old := loc
		loc += offsets[loc]
		if offsets[old] >= 3 {
			offsets[old]--
		} else {
			offsets[old]++
		}
		steps++
	}

	println(steps)
}
