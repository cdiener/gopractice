// Simple quiz program

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func getAnswer(record []string, stdin *bufio.Reader, out chan string) {
	fmt.Print(record[0], ": ")
	answer, _ := stdin.ReadString('\n')
	out <- strings.Trim(answer, "\n ")
	return
}

func main() {
	var limit int
	flag.IntVar(&limit, "time", 30, "the time limit in seconds")
	var shuffle bool
	flag.BoolVar(&shuffle, "shuffle", false, "whether to shuffle the questions")
	flag.Parse()
	path := strings.Join(flag.Args(), "")

	file, error := os.Open(path)
	if error != nil {
		fmt.Println("The file does not exist!")
		os.Exit(1)
	}

	r := csv.NewReader(file)
	stdin := bufio.NewReader(os.Stdin)
	correct := 0
	records, _ := r.ReadAll()
	total := len(records)
	timer := make(chan string, 1)

	var order []int
	if shuffle {
		order = rand.Perm(total)
	} else {
		order = make([]int, total)
		for i := range order {
			order[i] = i
		}
	}

	fmt.Print("Press Enter to start...")
	stdin.ReadString('\n')
	go func() {
		time.Sleep(time.Duration(limit) * time.Second)
		timer <- "done"
	}()

	for _, i := range order {
		record := records[i]
		answerChan := make(chan string, 1)
		go getAnswer(record, stdin, answerChan)
		select {
		case <-timer:
			fmt.Println("Time is up!")
			fmt.Printf("You got %d/%d correct!\n", correct, total)
			os.Exit(0)
		case answer := <-answerChan:
			if strings.ToLower(answer) == strings.ToLower(record[1]) {
				correct++
			}
		}
	}

	fmt.Printf("You got %d/%d correct!\n", correct, total)
}
