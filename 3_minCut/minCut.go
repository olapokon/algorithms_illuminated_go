package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// MinCut identifies the number of crossing edges for the minimum cat of a graph.
// parallel edges are allowed
func MinCut(graph map[int][]int) int {
	var edgeNum, randA, randB int
	for _, v := range graph {
		edgeNum += len(v)
	}
	edgeNum /= 2
	for len(graph) > 2 {
		// select an edge at random
		rand.Seed(time.Now().UnixNano())
		randEdge := rand.Intn(edgeNum)
		for k, v := range graph {
			if len(v) <= randEdge {
				randEdge -= len(v)
			} else {
				randA = k
				randB = v[randEdge]
			}
		}

		// merge nodes
		for _, node := range graph[randA] {
			// do not copy the edge if it connects to the node with which the merge is taking place,
			// to avoid self-loops
			if node == randB {
				continue
			}
			graph[randB] = append(graph[randB], node)
		}

		edgeNum--
		delete(graph, randA)
		// delete the merged node from the edges array of the node into which it is merged
		var i int
		for i < len(graph[randB]) {
			for graph[randB][i] == randA {
				graph[randB][i] = graph[randB][len(graph[randB])-1]
				graph[randB] = graph[randB][:len(graph[randB])-1]
				if i >= len(graph[randB]) {
					break
				}
			}
			i++
		}
		// go through all values in the graph and replace references to the merged node
		for key, val := range graph {
			if key == randB {
				continue
			}
			for i, v := range val {
				if v == randA {
					graph[key][i] = randB
				}
			}
		}
	}

	// return crossing edges
	var crossingEdges int
	for _, val := range graph {
		crossingEdges = len(val)
	}
	return crossingEdges
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
	graph := make(map[int][]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		// line := strings.Split(scanner.Text(), " ")
		line := strings.Split(scanner.Text(), "\t")
		nums := make([]int, len(line))
		for i := range line {
			if line[i] == "" {
				continue
			}
			nums[i], err = strconv.Atoi(line[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		graph[nums[0]] = nums[1:]
	}
	fmt.Printf("minimum cut crossing edges = %d\n", MinCut(graph))
}
