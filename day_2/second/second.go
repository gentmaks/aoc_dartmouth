//Package second
package second

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func SolveSecond () {
	path := "./day_2/first/input.txt"
	f, err := os.Open(path)
	res := 0
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}
	var tokens []string
	if line != "" {
		tokens = strings.Split(line, ",")
	}
	for _, token := range(tokens) {
		bounds := strings.Split(token, "-")
		first, err1 := strconv.Atoi(bounds[0])
		second, err2 := strconv.Atoi(bounds[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Could not convert the strings into ints")
			os.Exit(1)
		}
		for i := first; i <= second; i ++ {
			ok, length := validateSecond(strconv.Itoa(i))
			if ok {
				fmt.Println("The invalid number is: ", i, len(strconv.Itoa(i))/length)
				res += i
			}
		}
	}
	fmt.Println("The result is: ", res)
}

func validateID (s string) bool {
	if len(s) & 1 == 1{
		return false	
	}
	thresh := len(s) / 2 
	firstHalf := s[:thresh]
	secondHalf := s[thresh:]
	return firstHalf == secondHalf
}

func validateSecond(s string) (bool, int) {
	thresh := len(s) / 2
	for i := 1; i <= thresh; i++ {
		substr := s[:i]
		itr := i
		for itr < len(s) {
			if itr + i <= len(s) && s[itr:itr + i] == substr {
				if itr + i == len(s) {
					return true, i
				}
				itr += i
			} else {
				break
			}
		}
	}
	return false, -1
}
