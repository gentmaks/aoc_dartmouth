// Package first
package first

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gentmaks/aoc_dartmouth/internals/models"
)

func SolveFirst() {
	h := &models.FloatHeap{}
	populateDistances(h)
	heap.Init(h)
}

func populateDistances(h *models.FloatHeap) {
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
		}
	}
}
