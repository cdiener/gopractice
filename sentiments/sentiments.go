// A simple sentiment analyzer
package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Check whether an element can be found in an slice
func insert(s string, sentiment int, m map[string]int) bool {
	_, ok := m[s]
	if ok {
		return false
	}

	m[s] = sentiment

	return true
}

// Build the sentiment map
func buildMap(negativeFile string, positiveFile string) map[string]int {
	file, err := os.Open(negativeFile)
	defer file.Close()

	if err != nil {
		panic("could not read negative examples!")
	}
	var sentiments = make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		insert(strings.ToLower(scanner.Text()), -1, sentiments)
	}

	file, err = os.Open(positiveFile)
	defer file.Close()

	if err != nil {
		panic("could not read positive examples!")
	}
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		insert(strings.ToLower(scanner.Text()), 1, sentiments)
	}

	return sentiments
}

// Calculate the sentiment of a text
// a positive output means the text is positive etc.
func sentiment(text string, smap map[string]int) int {
	var sen = 0

	var val int
	var ok bool
	for _, word := range strings.Fields(text) {
		val, ok = smap[strings.ToLower(word)]
		if ok {
			sen += val
		}
	}

	return sen
}

func main() {
	posPtr := flag.String("pos", "positive.txt", "positive examples")
	negPtr := flag.String("neg", "negative.txt", "negative examples")
	flag.Parse()

	if len(flag.Args()) == 0 {
		println("Please give something to analyze :O")
		os.Exit(1)
	}

	phrase := strings.Join(flag.Args(), " ")

	sentiments := buildMap(*negPtr, *posPtr)
	sen := sentiment(phrase, sentiments)

	if sen > 0 {
		color.Green("%s (score=%d)", phrase, sen)
	} else if sen < 0 {
		color.Red("%s (score=%d)", phrase, sen)
	} else {
		color.White("%s (score=%d)", phrase, sen)
	}

	os.Exit(0)
}
