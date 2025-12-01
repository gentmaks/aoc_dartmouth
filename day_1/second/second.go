//Package second
package second

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)
func SolveSecond() {
	path := "./day_1/first/input.txt"
	fmt.Println(path)
	init := 50
	res := 0
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		op := line[0]
		rest, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		for i := range rest {
			_ = i
			if string(op) == "L" {
				init --
			} else {
				init ++ 
			}
			init = ((init % 100) + 100) % 100
			if init == 0 {
				res++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The amount of times 0 is reached is: ", res)
}
