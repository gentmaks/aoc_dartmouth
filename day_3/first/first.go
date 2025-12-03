//Package first
package first

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

func SolveFirst() {
	res := 0
	path := "./day_3/first/input.txt"
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		var bankHigh int
		for i := 0; i < len(line) - 1; i++ {
			for j := i + 1; j < len(line); j++ {
				num, err := strconv.Atoi(string(line[i] + line[j]))
				if err != nil {
					fmt.Println("Error converting string to integer")
					os.Exit(1)
				}
				bankHigh = max(bankHigh, num)
			}
		}
		res += bankHigh
	}
	fmt.Println("The result of adding all the banks is: ", res)
}
