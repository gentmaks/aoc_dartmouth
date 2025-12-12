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

type Region struct {
	area     int
	presents []int
}

func SolveFirst() {
	presentMap, regions := parse()
	count := solve(presentMap, regions)
	// fmt.Printf("presentMap: %v, regions: %v\n", presentMap, regions)
	fmt.Println("The answer is: ", count)
}

func parse() (map[int]int, []Region) {
	filePath := "./day_12/first/input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error when opening the file: ", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatal("Error when closing the file: ", err)
		}
	}(f)
	presentMap := map[int]int{}
	regions := []Region{}
	var counter int
	var currArea int
	var presentId int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if counter%5 == 0 {
			presentId, err = strconv.Atoi(string(line[0]))
			if err != nil {
				log.Fatal("Error when converting string to integer: ", err)
			}
		}
		if line == "" {
			presentMap[presentId] = currArea
			currArea = 0
			if _, ok := presentMap[5]; ok {
				break
			}
		}
		for _, ch := range line {
			if ch == '#' {
				currArea++
			}
		}
		counter++
	}
	for scanner.Scan() {
		var presents []int
		line := scanner.Text()
		newRegion := Region{}
		parts := strings.Split(line, ":")
		width, err1 := strconv.Atoi(parts[0][:2])
		height, err2 := strconv.Atoi(parts[0][3:])
		if err1 != nil || err2 != nil {
			log.Fatal("Error when converting string to integer: ", err1)
		}
		newRegion.area = width * height
		presentTokens := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, tok := range presentTokens {
			intTok, err := strconv.Atoi(tok)
			if err != nil {
				log.Fatal("Error when converting string to integer: ", err1)
			}
			presents = append(presents, intTok)
		}
		newRegion.presents = presents
		regions = append(regions, newRegion)
	}
	return presentMap, regions
}

func solve(presentMap map[int]int, regions []Region) int {
	count := 0
	for _, r := range regions {
		var currArea int
		for i, p := range r.presents {
			currArea += presentMap[i] * p
		}
		if currArea <= r.area {
			count++
		}
	}
	return count
}
