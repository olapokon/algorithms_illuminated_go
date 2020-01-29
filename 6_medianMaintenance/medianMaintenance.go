package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type minHeap struct {
	intHeap
}

type maxHeap struct {
	intHeap
}

func (h maxHeap) Less(i, j int) bool { return h.intHeap[i] > h.intHeap[j] }

// Medians returns the last four digits of the sum of medians of a stream of numbers,
// adding up the new median as a new number is being read from the stream.
func Medians(nums []int) int {
	var median, medianSum int
	minHeap := &minHeap{}
	maxHeap := &maxHeap{}

	for _, n := range nums {
		if maxHeap.Len() > 0 && n < (*maxHeap).intHeap[0] {
			heap.Push(maxHeap, n)
		} else if minHeap.Len() > 0 && n > (*minHeap).intHeap[0] {
			heap.Push(minHeap, n)
		} else {
			heap.Push(maxHeap, n)
		}

		// rebalance if the two heaps become uneven
		if minHeap.Len() > (maxHeap.Len() + 1) {
			temp := heap.Pop(minHeap)
			heap.Push(maxHeap, temp)
		}
		if maxHeap.Len() > (minHeap.Len() + 1) {
			temp := heap.Pop(maxHeap)
			heap.Push(minHeap, temp)
		}

		if minHeap.Len() > maxHeap.Len() {
			median = (*minHeap).intHeap[0]
		} else {
			median = (*maxHeap).intHeap[0]
		}
		medianSum += median
	}
	return medianSum % 10000
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

	var nums [10008]int
	var i int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = num
		i++
	}
	fmt.Printf("sum of medians mod 10000: = %d\n", Medians(nums[:i]))
}
