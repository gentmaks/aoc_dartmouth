// Package first
package first

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SolveFirst() {
	data := parse()
	graph := buildGraph(data)
	fmt.Println(graph)
	visited["svr"] = struct{}{}
	dfs(graph, "svr", 2)
	fmt.Println(res)
}

var (
	res     = 0
	visited = map[string]struct{}{}
)

func parse() [][]string {
	data := [][]string{}
	filePath := "./day_11/first/demo_input.txt"
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

func dfs(graph map[string][]string, curr string, rem int) {
	fmt.Println("curr: ", curr)
	if curr == "out" {
		if rem == 0 {
			res++
		}
		return
	}
	for _, nei := range graph[curr] {
		_, ok := visited[nei]
		if ok {
			continue
		}
		if nei == "fft" || nei == "dac" {
			visited[nei] = struct{}{}
			dfs(graph, nei, rem-1)
		} else {
			visited[nei] = struct{}{}
			dfs(graph, nei, rem)
		}
	}
}
