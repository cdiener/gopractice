// finds a number in a list of numbers

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find(nums []int64, what int64) bool {
	left := 0
	right := len(nums) - 1
	mid := int(right / 2)

	for right > left {
		if nums[mid] == what {
			return true
		}
		if nums[mid] > what {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = int(left/2 + right/2)
	}

	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need exactly one argument!")
		os.Exit(1)
	}
	what, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		fmt.Println("Argument is not a whole number!")
		os.Exit(1)
	}
	var nums []int64
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	w := bufio.NewScanner(strings.NewReader(input))
	w.Split(bufio.ScanWords)

	var i int64
	for w.Scan() {
		i, err = strconv.ParseInt(w.Text(), 0, 0)
		if err != nil {
			panic(fmt.Sprintf("%s is not a number", w.Text()))
		}
		nums = append(nums, i)
	}

	found := find(nums, what)
	if found {
		fmt.Println("Found it!")
		os.Exit(0)
	} else {
		fmt.Println("Nope, not here :(")
		os.Exit(1)
	}
}
