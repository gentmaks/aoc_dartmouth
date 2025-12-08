// Package first
package first

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gentmaks/aoc_dartmouth/internals/models"
)

func SolveFirst() {
	h := &models.VectorHeap{}
	populateDistances(h)
	heap.Init(h)
	fmt.Println((*h)[0])
}

func populateDistances(h *models.VectorHeap) {
	coords := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		coords[i] = make([]int, 3)
	}
	filePath := "./day_8/first/input.txt"
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
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		for i, token := range tokens {
			value, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal("Error when converting string to integer: ", err)
			}
			coords[row][i] = value
		}
		row++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error when scanning the file: ", err)
	}
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			vec := models.Vector{
				Dist: getDistance(coords[i],
					coords[j]),
				Source: i,
				Target: j,
			}
			h.Push(vec)
		}
	}
}

func getDistance(a, b []int) float64 {
	dx := float64(a[0] - b[0])
	dy := float64(a[1] - b[1])
	dz := float64(a[2] - b[2])

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
