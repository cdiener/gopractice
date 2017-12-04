package main

import (
	"flag"
	"sort"
	"strings"
)

func validate(pass string) bool {
	words := strings.Fields(pass)
	counter := make(map[string]int)
	for _, word := range words {
		letters := strings.Split(word, "")
		sort.Strings(letters)
		sorted := strings.Join(letters, "")
		_, present := counter[sorted]
		if present {
			return false
		}
		counter[sorted] = 1
	}
	return true
}

func main() {
	flag.Parse()
	passphrases := strings.Split(flag.Args()[0], "\n")
	valid := 0
	for _, pass := range passphrases {
		if validate(pass) {
			valid++
		}
	}

	println(valid)
}
