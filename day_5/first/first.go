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

type Range struct {
	lower int
	upper int
}

func SolveFirst() {
	ranges, ids := parseFile()
	res := getValidCount(ranges, ids)
	fmt.Println("The count of valid ids is: ", res)
}

func parseFile() ([]Range, []int) {
	filePath := "./day_5/first/input.txt"
	ranges := []Range{}
	ids := []int{}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error when reading the input file: ", err)
	}
	scanner := bufio.NewScanner(f)
	var idOk bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			idOk = true
			continue
		}
		if !idOk {
			tokens := strings.Split(line, "-")
			first, err1 := strconv.Atoi(tokens[0])
			second, err2 := strconv.Atoi(tokens[1])
			if err1 != nil || err2 != nil {
				log.Fatal("Error when converting the strings to integers: ", err)
			}
			ranges = append(ranges, Range{lower: first, upper: second})
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Error when converting the strings to integers: ", err)
			}
			ids = append(ids, val)
		}
	}
	return ranges, ids
}

func getValidCount(ranges []Range, ids []int) int {
	var count int
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.lower && id <= r.upper {
				count++
				break
			}
		}
	}
	return count
}
