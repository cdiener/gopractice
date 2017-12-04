package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

func addArray(x [2]int, y [2]int) ([2]int, error) {
	if len(x) != len(y) {
		return x, errors.New("arrays must have same length")
	}

	for i := range x {
		x[i] += y[i]
	}
	return x, nil
}

func emptyMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	return m
}

func areaSum(m [][]int, i int, j int) int {
	sum := m[i-1][j-1] + m[i-1][j] + m[i-1][j+1] + m[i][j-1] + m[i][j+1] +
		m[i+1][j-1] + m[i+1][j] + m[i+1][j+1]
	return sum
}

func printMatrix(m [][]int) {
	for _, v := range m {
		for _, k := range v {
			fmt.Printf("%5d", k)
		}
		fmt.Println()
	}
}

func main() {
	DIRS := [4][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	flag.Parse()

	num, _ := strconv.Atoi(flag.Args()[0])
	dim := int(math.Sqrt(float64(num)))
	matrix := emptyMatrix(dim + 5)
	var loc [2]int
	loc[0] = (dim + 5) / 2
	loc[1] = (dim + 5) / 2
	steps := 1
	turn := 0
	matrix[loc[0]][loc[1]] = 1

	for i := 0; i < dim; i++ {
		for turns := 0; turns < 2; turns++ {
			for k := 0; k < steps; k++ {
				loc, _ = addArray(loc, DIRS[turn%4])
				asum := areaSum(matrix, loc[0], loc[1])
				if asum > num {
					fmt.Println(asum)
					os.Exit(0)
				}
				matrix[loc[0]][loc[1]] = asum
			}
			turn++
		}
		steps++
	}
}
