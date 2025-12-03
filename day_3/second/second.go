//Package second
// Took inspo from: https://www.reddit.com/r/adventofcode/comments/1pcxkif/2025_day_3_mega_tutorial/
package second

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"math"
)

const INF = 1 << 40

type Key struct {
	sequence string
	movesLeft int
}
func SolveSecond() {
	res := 0
	path := "./day_3/first/input.txt"
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		cache := make(map[Key]int)
		line = scanner.Text()
		bankHigh := dp(Key{sequence: line, movesLeft: 12}, cache)
		fmt.Printf("For the following sequence: %s we get the following num: %d\n", line, bankHigh)
		res += bankHigh
	}
	fmt.Println("The result of adding all the banks is: ", res)
}

func dp(key Key, cache map[Key]int) int {
	if key.movesLeft == 0 {
		return 0
	}
	if key.sequence == "" {
		return -INF
	}
	value, ok := cache[key]
	if ok {
		return value
	}
	firstNum, err := strconv.Atoi(string(key.sequence[0])) 
	if err != nil {
		fmt.Println("Error converting string to integer")
		os.Exit(1)
	}
	result := max(
		firstNum*exp(10, key.movesLeft-1) + dp(Key{sequence: key.sequence[1:], movesLeft: key.movesLeft - 1}, cache),
		dp(Key{sequence: key.sequence[1:], movesLeft: key.movesLeft}, cache),
	)
	cache[key] = result
	return  result
}

func exp(base int, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}
