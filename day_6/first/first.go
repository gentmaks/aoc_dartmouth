// Package first
package first

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolveFirst() {
	grid, op := parseFile()
	res := calculate(grid, op)
	fmt.Println("Result is: ", res)
}

func parseFile() ([][]int, map[int]int) {
	filePath := "./day_6/first/input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error when reading the input file: ", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatal("Error when closing file: ", err)
		}
	}(f)
	rows, cols := getDimensions(f)
	grid := make([][]int, rows-1)
	for i := 0; i < rows-1; i++ {
		grid[i] = make([]int, cols)
	}
	op := make(map[int]int)
	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatal("Error when reseting the file pointer: ", err)
	}
	scanner := bufio.NewScanner(f)
	var c int
	for scanner.Scan() {
		if c == rows-1 {
			line := strings.TrimSpace(scanner.Text())
			tokens := strings.Fields(line)
			for i, token := range tokens {
				if token == "+" {
					op[i] = 0
				} else {
					op[i] = 1
				}
			}
			break
		}
		line := strings.TrimSpace(scanner.Text())
		tokens := strings.Fields(line)
		for i, token := range tokens {
			val, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal("Error when converting string to integer: ", err)
			}
			grid[c][i] = val
		}
		c++
	}
	return grid, op
}

func getDimensions(f *os.File) (int, int) {
	var rows int
	var cols int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rows++
		line := strings.TrimSpace(scanner.Text())
		tokens := strings.Fields(line)
		cols = len(tokens)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Could not get the number of rows and columns from the file", err)
	}
	f.Truncate(0)
	return rows, cols
}

func calculate(grid [][]int, op map[int]int) int {
	var res int
	for colIdx := range grid[0] {
		var mult bool
		var curr int
		if op[colIdx] == 0 {
			curr = 0
		} else {
			curr = 1
			mult = true
		}
		for rowIdx := range grid {
			val := grid[rowIdx][colIdx]
			if mult {
				curr *= val
			} else {
				curr += val
			}
		}
		res += curr
	}
	return res
}
