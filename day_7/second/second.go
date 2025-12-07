// Package second
package second

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SolveSecond() {
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
	lines := make(map[int]int)
	for col := 0; col < 141; col++ {
		lines[col] = 0
	}
	lines[70] = 1
	for i, row := range grid[1:] {
		if i == 14 {
			fmt.Println(lines)
			currSum := 0
			for _, v := range lines {
				currSum += v
			}
			fmt.Println(currSum)
		}
		for j, val := range row {
			if val == "^" {
				count := lines[j]
				lines[j] = 0
				left := j - 1
				right := j + 1
				if left >= 0 {
					if val := lines[left]; val == 0 {
						lines[left] = count
					} else {
						lines[left] = lines[left] + count
					}
				}
				if right < len(grid[0]) {
					if val := lines[left]; val == 0 {
						lines[right] = count
					} else {
						lines[right] = lines[right] + count
					}
				}
			}
		}
		if i == 140 {
			fmt.Println("inside")
			for _, value := range lines {
				res += value
			}
		}
	}
	return res
}
