// Package second
package second

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type State struct {
	node     string
	fftFound bool
	dacFound bool
}

func SolveSecond() {
	data := parse()
	graph := buildGraph(data)
	// fmt.Println("Initial graph: ", graph)
	res := dfs(graph, State{node: "svr", fftFound: false, dacFound: false})
	fmt.Println(res)
}

var memo = map[State]int{}

func parse() [][]string {
	data := [][]string{}
	filePath := "./day_11/first/input.txt"
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
		curr := []string{}
		line := scanner.Text()
		tokens := strings.Split(line, ":")
		curr = append(curr, tokens[0])
		dst := strings.Split(strings.TrimSpace(tokens[1]), " ")
		curr = append(curr, dst...)
		data = append(data, curr)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error when reading from the file: ", err)
	}
	return data
}

func buildGraph(data [][]string) map[string][]string {
	g := map[string][]string{}
	for _, row := range data {
		for i := 1; i < len(row); i++ {
			g[row[0]] = append(g[row[0]], row[i])
		}
	}
	return g
}

func dfs(graph map[string][]string, s State) int {
	fmt.Println(s)
	fmt.Println("graph: ", graph["srv"])
	if v, ok := memo[s]; ok {
		return v
	}
	if s.node == "out" {
		if s.fftFound && s.dacFound {
			return 1
		}
		return 0
	}
	tally := 0
	for _, nei := range graph[s.node] {
		next := State{
			node:     nei,
			fftFound: s.fftFound,
			dacFound: s.dacFound,
		}
		if nei == "fft" {
			next.fftFound = true
		}
		if nei == "dac" {
			next.dacFound = true
		}
		tally += dfs(graph, next)
	}
	memo[s] = tally
	return tally
}
