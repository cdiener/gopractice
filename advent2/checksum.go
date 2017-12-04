package main

import (
	"flag"
	"strconv"
	"strings"
)

// MinMaxDiff calculates the difference between max and min
// for a slice of values
func MinMaxDiff(vals []string) int {
	min, _ := strconv.Atoi(vals[0])
	max := min
	for _, v := range vals {
		d, err := strconv.Atoi(v)
		if err != nil {
			panic("could not parse number!")
		}
		if int(d) < min {
			min = int(d)
		} else if int(d) > max {
			max = int(d)
		}
	}

	return max - min
}

func evenDivivion(vals []string) int {
	for i, vi := range vals {
		di, _ := strconv.Atoi(vi)
		for j, vj := range vals {
			if i == j {
				continue
			}
			dj, _ := strconv.Atoi(vj)
			if di%dj == 0 {
				return di / dj
			}
			if dj%di == 0 {
				return dj / di
			}
		}
	}
	return -1
}

func main() {
	flag.Parse()
	input := strings.Join(flag.Args(), "")
	nums := strings.Split(input, "\n")
	sum := 0

	for _, s := range nums {
		vals := strings.Fields(s)
		sum += evenDivivion(vals)
	}

	println(sum)
}
