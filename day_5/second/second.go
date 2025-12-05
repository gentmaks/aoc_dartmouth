//Package second
package second

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"sort"
)

type Range struct {
	lower int; 
	upper int
}

func SolveSecond() {
	ranges, _ := parseFile()
	res := getTotalValid(ranges)
	fmt.Println("The count of total valid ids is: ", res)
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
	for _, id := range(ids) {
		for _, r := range(ranges){
			if id >= r.lower && id <= r.upper {
				count++
				break
			}
		}
	}
	return count
}

func getTotalValid(ranges []Range) int {
	sort.Slice(ranges, func(i, j int) bool {
        return ranges[i].lower < ranges[j].lower
    })
    newRanges := []Range{}

    for _, r := range ranges {
        merged := false

        for i := range newRanges {
            if r.lower <= newRanges[i].upper && r.upper >= newRanges[i].lower {
                newRanges[i].lower = min(r.lower, newRanges[i].lower)
                newRanges[i].upper = max(r.upper, newRanges[i].upper)
                merged = true
                break
            }
        }

        if !merged {
            newRanges = append(newRanges, r)
        }
    }

    res := 0
    for _, r := range newRanges {
        res += (r.upper - r.lower + 1)
    }
    return res
}
