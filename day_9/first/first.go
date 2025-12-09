// Package first
package first

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func SolveFirst() {
	fmt.Println("hi")
	coords := Parse()
	res := Solve(coords)
	fmt.Println(res)
}

func Parse() []Coord {
	coords := []Coord{}
	filePath := "./day_9/first/input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error when opening the file: ", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatal("Error when closing the file: ", err)
		}
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		xVal, err1 := strconv.Atoi(tokens[0])
		yVal, err2 := strconv.Atoi(tokens[1])
		if err1 != nil || err2 != nil {
			log.Fatal("Error when converting the strings to integers: ", err1)
		}
		coords = append(coords, Coord{x: xVal, y: yVal})
	}
	return coords
}

func Solve(coords []Coord) int {
	maxArea := 0
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			first, second := coords[i], coords[j]
			currArea := (math.Abs(float64(first.x)-float64(second.x)) + 1) * math.Abs(float64(first.y)-float64(second.y)+1)
			maxArea = max(maxArea, int(currArea))
		}
	}
	return maxArea
}

