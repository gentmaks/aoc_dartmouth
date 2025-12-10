// Package first
package first

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Button struct {
	lights []int
}

func SolveFirst() {
	table := parse()
	res := 0
	for i := range table {
		buttons := []Button{}
		bits := getOpenBits(table[i][0])
		for _, seq := range table[i][1 : len(table[i])-1] {
			butt := getButton(seq)
			buttons = append(buttons, butt)
		}
		comb := generateComb(buttons)
		shortest := getShortest(bits, comb)
		res += shortest
		fmt.Printf("Bracket: %v, Bits: %v, Index: %d, Shortest: %d\n", table[i][0], bits, i, shortest)
		fmt.Println("Bracket and bits: ", table[i][0], bits, i)
	}
	fmt.Println(res)
}

func parse() [][]string {
	table := [][]string{}
	filePath := "day_10/first/input.txt"
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
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		table = append(table, tokens)
	}
	return table
}

func getButton(buttonString string) Button {
	lights := []int{}
	buttonString = buttonString[1 : len(buttonString)-1]
	nums := strings.Split(buttonString, ",")
	for _, n := range nums {
		intNum, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("Error when casting string into integer: ", err)
		}
		lights = append(lights, intNum)
	}
	return Button{lights: lights}
}

func generateComb(buttonList []Button) [][]Button {
	comb := [][]Button{}
	currComb := []Button{}
	dp(0, buttonList, &comb, &currComb)
	sort.Slice(comb, func(i int, j int) bool {
		return len(comb[i]) < len(comb[j])
	})
	return comb
}

func dp(i int, buttonList []Button, comb *[][]Button, currComb *[]Button) {
	if i == len(buttonList) {
		*comb = append(*comb, cloneButtons(*currComb))
		return
	}

	*currComb = append(*currComb, buttonList[i])
	dp(i+1, buttonList, comb, currComb)
	*currComb = (*currComb)[:len(*currComb)-1]
	dp(i+1, buttonList, comb, currComb)
}

func cloneButtons(src []Button) []Button {
	dst := make([]Button, len(src))
	for i := range src {
		newLights := make([]int, len(src[i].lights))
		copy(newLights, src[i].lights)
		dst[i] = Button{lights: newLights}
	}
	return dst
}

func getOpenBits(bracket string) []int {
	bits := []int{}
	for i := 1; i < len(bracket)-1; i++ {
		if string(bracket[i]) == "#" {
			bits = append(bits, 1)
		} else {
			bits = append(bits, 0)
		}
	}
	return bits
}

func getShortest(dstBits []int, comb [][]Button) int {
	for _, c := range comb {
		srcBits := make([]int, len(dstBits))
		for _, b := range c {
			for _, l := range b.lights {
				srcBits[l] = 1 - srcBits[l]
			}
		}
		if compareSlices(srcBits, dstBits) {
			return len(c)
		}
	}
	return -1
}

func compareSlices(a []int, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
