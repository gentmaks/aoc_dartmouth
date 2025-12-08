// Package first
package first

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gentmaks/aoc_dartmouth/internals/models"
)

func SolveFirst() {
	h := &models.VectorHeap{}
	uf := models.UFInit(1000)
	populateDistances(h)
	fmt.Println("Length of the heap is: ", len(*h))
	heap.Init(h)
	fmt.Println("The elements in the heap are: ", *h)
	count := 0
	for count < 1000 {
		v := heap.Pop(h).(models.Vector)
		d := v.Dist
		s := v.Source
		t := v.Target
		fmt.Println("The smallest distance currently is: ", d)
		uf.Union(s, t)
		count++
	}
	fmt.Println("Connected component counts: ", uf.GetConnCompCount())
	fmt.Println("The answer is: ", getAnswer(uf))
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

func getAnswer(uf *models.UnionFind) int {
	sizes := uf.GetSizeArray()
	sort.Ints(sizes)
	last := len(sizes) - 1
	fmt.Println("The sizes array is: ", sizes)
	fmt.Println("The size of the biggest conn comp is: ", sizes[last])
	return (sizes[last] * sizes[last-1] * sizes[last-2])
}
