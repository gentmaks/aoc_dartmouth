// Package second
package second

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Offset struct {
	x int
	y int
}

func SolveSecond() {
	path := "./day_4/first/input.txt"
	res := 0
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatalln("Could not close file", err)
		}
	}(f)
	rows, cols := getDimensions(f)
	grid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]string, cols)
	}
	ok := populateGrid(f, &grid)
	if !ok {
		log.Fatalln("Could not populate grid")
	}
	for {
		if curr := getValidCount(&grid, rows, cols); curr > 0 {
			res += curr
		} else {
			break
		}
	}
	fmt.Println(res)
}

func checkValid(r int, c int, numRows int, numCols int, grid [][]string) bool {
	var neiCount int
	directions := []Offset{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}
	for _, dir := range directions {
		newR := r + dir.x
		newC := c + dir.y
		if newR < 0 || newR >= numRows || newC < 0 || newC >= numCols || grid[newR][newC] != "@" {
			continue
		}
		neiCount++
	}
	return neiCount < 4
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

func getValidCount(grid *[][]string, numRows int, numCols int) int {
	res := 0
	for i, row := range *grid {
		for j, col := range row {
			if col == "@" && checkValid(i, j, numRows, numCols, *grid) {
				res++
				(*grid)[i][j] = "x"
			}
		}
	}
	return res
}

func populateGrid(f *os.File, grid *[][]string) bool {
	_, err := f.Seek(0, 0)
	if err != nil {
		fmt.Println("failed to seek:", err)
		return false
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

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%q", col)
		}
		fmt.Printf("\n")
	}
}
