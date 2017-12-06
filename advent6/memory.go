package main

import (
	"flag"
	"strconv"
	"strings"
)

func maxIndex(banks []int) int {
	maxi := 0
	max := banks[0]
	for i, v := range banks {
		if v > max {
			maxi = i
			max = v
		}
	}
	return maxi
}

func redistribute(bankp *[]int) {
	banks := *bankp
	n := len(banks)
	start := maxIndex(banks)
	load := banks[start]
	banks[start] = 0
	for {
		start = (start + 1) % n
		banks[start]++
		load--
		if load == 0 {
			break
		}
	}
}

func hash(banks []int) string {
	strs := make([]string, len(banks))
	for i := range banks {
		strs[i] = strconv.Itoa(banks[i])
	}
	return strings.Join(strs, "|")
}

func main() {
	flag.Parse()
	input := flag.Args()
	banks := make([]int, len(input))
	for i := range banks {
		banks[i], _ = strconv.Atoi(input[i])
	}

	known := make(map[string]int)
	known[hash(banks)] = 0
	runs := 0
	for {
		runs++
		redistribute(&banks)
		h := hash(banks)
		idx, present := known[h]
		if present {
			println(runs - idx)
			break
		} else {
			known[h] = runs
		}
	}

}
