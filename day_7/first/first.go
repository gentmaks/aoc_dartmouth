// Package first
package first

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SolveFirst() {
	filePath := "./day_7/first/input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error in opening the file: ", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatal("Error in closing the file: ", err)
		}
	}(f)
	rows, cols := getDimensions(f)
	grid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]string, cols)
	}
	ok := populateGrid(f, &grid)
	if !ok {
		log.Fatal("Error when populating the grid")
	}
	res := calculate(grid)
	fmt.Println("The result is: ", res)
}

func getDimensions(f *os.File) (int, int) {
	var rows int
	var cols int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rows++
		line := scanner.Text()
		cols = len(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Could not get the number of rows and columns from the file", err)
	}
	return rows, cols
}

func populateGrid(f *os.File, grid *[][]string) bool {
	_, err := f.Seek(0, 0)
	if err != nil {
		log.Fatal("Error when resetting the file pointer: ", err)
	}
	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, char := range line {
			(*grid)[row][col] = string(char)
		}
		row++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Could not populate grid")
		return false
	}
	return true
}

func calculate(grid [][]string) int {
	res := 0
	lines := make(map[int]struct{})
	lines[70] = struct{}{}
	for i, row := range grid[1:] {
		for j, val := range row {
			if val == "^" {
				if _, ok := lines[j]; ok {
					res++
					delete(lines, j)
				}
				newRow := i + 1
				if newRow >= len(grid) {
					return res
				}
				left := j - 1
				right := j + 1
				if left >= 0 {
					lines[left] = struct{}{}
				}
				if right < len(grid[0]) {
					lines[right] = struct{}{}
				}
			}
		}
	}
	return res
}
