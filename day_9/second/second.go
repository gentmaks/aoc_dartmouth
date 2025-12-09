// Package second
package second

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Rect struct {
	p1 Coord
	p2 Coord
}

var Boundary = make(map[Coord]struct{})

func SolveSecond() {
	coords := Parse()
	rectangles := getRectangles(coords)
	sort.Slice(rectangles, func(i, j int) bool {
		return getRectArea(rectangles[i]) > getRectArea(rectangles[j])
	})
	for i := 0; i < len(coords)-1; i++ {
		getLinePoints(coords[i], coords[i+1])
	}
	getLinePoints(coords[0], coords[len(coords)-1])
	for _, rect := range rectangles {
		coords := getPerimPoints(rect.p1, rect.p2)
		area := getRectArea(rect)
		if !detectCollision(coords) {
			fmt.Println(area)
			return
		}
	}
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

func getLinePoints(a Coord, b Coord) {
	for i := min(a.x, b.x); i <= max(a.x, b.x); i++ {
		c := Coord{x: i, y: a.y}
		_, ok := Boundary[c]
		if !ok {
			Boundary[c] = struct{}{}
		}
	}
	for j := min(a.y, b.y); j <= max(a.y, b.y); j++ {
		c := Coord{x: a.x, y: j}
		_, ok := Boundary[c]
		if !ok {
			Boundary[c] = struct{}{}
		}
	}
}

func getPerimPoints(a Coord, b Coord) []Coord {
	coords := []Coord{}
	if math.Abs(float64(a.x)-float64(b.x)) <= 2 || math.Abs(float64(a.y)-float64(b.y)) <= 2 {
		return nil
	}
	minX, maxX := min(a.x, b.x), max(a.x, b.x)
	minY, maxY := min(a.y, b.y), max(a.y, b.y)
	minX++
	maxX--
	minY++
	maxY--
	for i := minX; i <= maxX; i++ {
		coords = append(coords, Coord{x: i, y: minY})
		coords = append(coords, Coord{x: i, y: maxY})
	}
	for j := minY; j <= maxY; j++ {
		coords = append(coords, Coord{x: minX, y: j})
		coords = append(coords, Coord{x: maxX, y: j})
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

func getRectangles(coords []Coord) []Rect {
	rectangles := []Rect{}
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			first, second := coords[i], coords[j]
			rectangles = append(rectangles, Rect{p1: first, p2: second})
		}
	}
	return rectangles
}

func getRectArea(rect Rect) int {
	first, second := rect.p1, rect.p2
	currArea := (math.Abs(float64(first.x)-float64(second.x)) + 1) * math.Abs(float64(first.y)-float64(second.y)+1)
	return int(currArea)
}

func detectCollision(coords []Coord) bool {
	for _, coord := range coords {
		_, ok := Boundary[coord]
		if ok {
			fmt.Println("running")
			return true
		}
	}
	return false
}
