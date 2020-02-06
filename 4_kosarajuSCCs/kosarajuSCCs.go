package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type vertex struct {
	f        int
	explored bool
	scc      int
}

var curLabel, numSCC int

// Input: directed acyclic graph G = (V, E) in
// adjacency-list representation.
// Postcondition: the f -values of vertices constitute a
// topological ordering of G.
func topoSort(edges map[int][]int, vertices map[int]*vertex) {
	curLabel = len(vertices)
	for k, v := range vertices {
		if !v.explored {
			dFSTopo(edges, vertices, k)
		}
	}
}

// Input: graph G = (V, E) in adjacency-list
// representation, and a vertex s in V .
// Postcondition: every vertex reachable from s is
// marked as “explored” and has an assigned f -value.
func dFSTopo(edges map[int][]int, vertices map[int]*vertex, s int) {
	vertices[s].explored = true
	for _, v := range edges[s] {
		if !vertices[v].explored {
			dFSTopo(edges, vertices, v)
		}
	}
	vertices[s].f = curLabel
	curLabel--
}

// Input: directed graph G = (V, E) in adjacency-list
// representation, and a vertex s in V .
// Postcondition: every vertex reachable from s is
// marked as “explored” and has an assigned scc-value.
func dFSSCC(edges map[int][]int, vertices map[int]*vertex, s int) {
	vertices[s].explored = true

	// scc(s) := numSCC // global variable above
	vertices[s].scc = numSCC

	for _, v := range edges[s] {
		if !vertices[v].explored {
			dFSSCC(edges, vertices, v)
		}
	}
}

// Kosaraju detects the strongly connected components of a directed graph.
//
// Input: directed graph G = (V, E) in adjacency-list
// representation, with V = {1, 2, 3, . . . , n}
// Postcondition: for every v, w in V , scc(v) = scc(w)
// if and only if v, w are in the same SCC of G.
func Kosaraju(edges, edgesRev map[int][]int, vertices map[int]*vertex) []int {

	// first pass of depth-first search
	// (computes f (v)’s, the magical ordering)
	topoSort(edgesRev, vertices)

	// second pass of depth-first search
	// (finds SCCs in reverse topological order)

	// create an array with the elements in increasing f(v) order
	order := make([]int, len(vertices))
	for k, v := range vertices {
		order[v.f-1] = k
	}

	// mark all vertices of G as unexplored
	for k := range vertices {
		vertices[k].explored = false
	}

	for _, v := range order {
		if !vertices[v].explored {
			numSCC++
			dFSSCC(edges, vertices, v)
		}
	}

	// calculate SCC sizes
	sccs := make(map[int]int)
	for _, v := range vertices {
		if _, ok := sccs[v.scc]; ok {
			sccs[v.scc]++
		} else {
			sccs[v.scc] = 1
		}
	}

	// enter sizes in an array and sort
	sccSizes := make([]int, len(sccs))
	for _, v := range sccs {
		sccSizes = append(sccSizes, v)
	}
	sort.Ints(sccSizes)
	return sccSizes
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

	edges := make(map[int][]int)
	vertices := make(map[int]*vertex)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		line := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		nums := make([]int, 2)
		for i := range line {
			nums[i], err = strconv.Atoi(line[i])
			if err != nil {
				log.Fatal(err)
			}
		}

		// enter vertices in the vertices map
		for _, v := range nums {
			if _, ok := vertices[v]; !ok {
				vertices[v] = &vertex{0, false, 0}
			}
		}

		// enter relationships between vertices in the edges map
		if _, ok := edges[nums[0]]; ok {
			edges[nums[0]] = append(edges[nums[0]], nums[1])
		} else {
			edges[nums[0]] = nums[1:]
		}
	}
	// create reversed edges map
	edgesRev := make(map[int][]int)
	for k, v := range edges {
		for _, e := range v {
			if _, ok := edgesRev[e]; ok {
				edgesRev[e] = append(edgesRev[e], k)
			} else {
				edgesRev[e] = []int{k}
			}
		}
	}

	sccSizes := Kosaraju(edges, edgesRev, vertices)
	fmt.Println(sccSizes)
}
