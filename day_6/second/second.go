// Package second
package second

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolveSecond() {
	grid := parseFile()
	res := calculate(grid)
	fmt.Println("Result is: ", res)
	// res := calculate(grid, op)
	// fmt.Println("Result is: ", res)
}

func parseFile() [][]string {
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
	rows := 5
	cols := 3752
	grid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]string, cols)
	}
	scanner := bufio.NewScanner(f)
	var c int
	for scanner.Scan() {
		line := scanner.Text()
		for idx, char := range line {
			grid[c][idx] = string(char)
		}
		c++
	}
	return grid
}

// From looking at the file we know dimensions are --> rows: 5, cols: 3752
// func getDimensions(f *os.File) (int, int) {
// 	var rows int
// 	var cols int
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		rows++
// 		line := strings.TrimSpace(scanner.Text())
// 		tokens := strings.Fields(line)
// 		cols = len(tokens)
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal("Could not get the number of rows and columns from the file", err)
// 	}
// 	f.Truncate(0)
// 	return rows, cols
// }

// func calculate(grid [][]string, op map[int]int) int {
// 	var res int
// 	for colIdx := range grid[0] {
// 		var mult bool
// 		var curr int
// 		if op[colIdx] == 0 {
// 			curr = 0
// 		} else {
// 			curr = 1
// 			mult = true
// 		}
// 		for rowIdx := range grid {
// 			val := grid[rowIdx][colIdx]
// 			if mult {
// 				curr *= val
// 			} else {
// 				curr += val
// 			}
// 		}
// 		res += curr
// 	}
// 	return res
// }

func calculate(grid [][]string) int {
	res := 0
	op := "*"
	currTally := 1
	var bp bool
	for colIdx := range grid[0] {
		currNum := grid[0][colIdx] + grid[1][colIdx] + grid[2][colIdx] + grid[3][colIdx]
		if currNum == "    " {
			bp = true
			continue
		}
		if bp {
			fmt.Printf("ColIdx: %d, currTally: %d, res: %d\n", colIdx, currTally, res)
			op = grid[4][colIdx]
			bp = false
			res += currTally
			if op == "+" {
				currTally = 0
			} else {
				currTally = 1
			}
		}
		val, err := strconv.Atoi(strings.TrimSpace(currNum))
		if err != nil {
			log.Fatal("Error converting string to integer: ", err)
		}
		if op == "*" {
			if val != 0 {
				currTally *= val
			}
		} else {
			currTally += val
		}
	}
	res += currTally
	return res
}
