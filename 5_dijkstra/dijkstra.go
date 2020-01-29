package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Head represents the head vertex of an edge, along with its length.
type Head struct {
	head, length int
}

// Dijkstra takes a directed graph as input, and returns the shortest
// paths from the source vertex 1 to every other vertex in the graph.
func Dijkstra(edges map[int][]Head) map[int]int {
	// initialize a table with all lenghts set to infinity/max integer
	lengths := make(map[int]int)
	for k := range edges {
		lengths[k] = math.MaxInt64
	}
	// initialize a table that holds explored vertices
	explored := make(map[int]bool, len(edges))
	// using vertex 1 as the source
	lengths[1] = 0
	explored[1] = true

	// compare edges with an explored tail and unexplored head
	var head, length int
	for len(explored) < len(edges) {
		length = math.MaxInt64
		for k := range explored {
			for _, e := range edges[k] {
				if _, ok := explored[e.head]; !ok {
					// it is an unexplored edge
					if lengths[k]+e.length < length {
						length = lengths[k] + e.length
						head = e.head
					}
				}
			}
		}
		// mark the head node with the shortest length as explored
		explored[head] = true
		lengths[head] = length
	}
	return lengths
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("file argument missing")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	edges := make(map[int][]Head)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		tail, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		edges[tail] = make([]Head, len(line)-1)
		for i := 1; i < len(line); i++ {
			if line[i] == "" {
				continue
			}
			h := Head{}
			h.head, err = strconv.Atoi(line[i][:strings.Index(line[i], ",")])
			if err != nil {
				log.Fatal(err)
			}
			h.length, err = strconv.Atoi(line[i][strings.Index(line[i], ",")+1:])
			if err != nil {
				log.Fatal(err)
			}
			edges[tail][i-1] = h
		}
	}
	lengths := Dijkstra(edges)
	fmt.Printf("vertex 7: %d\n", lengths[7])
	fmt.Printf("vertex 37: %d\n", lengths[37])
	fmt.Printf("vertex 59: %d\n", lengths[59])
	fmt.Printf("vertex 82: %d\n", lengths[82])
	fmt.Printf("vertex 99: %d\n", lengths[99])
	fmt.Printf("vertex 115: %d\n", lengths[115])
	fmt.Printf("vertex 133: %d\n", lengths[133])
	fmt.Printf("vertex 165: %d\n", lengths[165])
	fmt.Printf("vertex 188: %d\n", lengths[188])
	fmt.Printf("vertex 197: %d\n", lengths[197])
}
