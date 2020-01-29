package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// TwoSUM takes as input a range of numbers and a set of integers, and returns
// the number of numbers t in the given range, for which 2 distinct numbers x and y exist
// in the set, such that x + y = t
func TwoSUM(start, end int, nums map[int]int, c chan int) {
	var count int
	for t := start; t <= end; t++ {
		for x := range nums {
			y, ok := nums[t-x]
			if ok && y != x {
				count++
				fmt.Printf("t = %d, x = %d, y = %d\n", t, x, y)
				break
			}
		}
	}
	c <- count
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

	nums := make(map[int]int)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums[num] = num
	}

	c := make(chan int)
	for i := -10000; i <= 8000; i += 2000 {
		go TwoSUM(i, i+2000, nums, c)
	}
	c1, c2, c3, c4, c5, c6, c7, c8, c9, c10 := <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c
	count := c1 + c2 + c3 + c4 + c5 + c6 + c7 + c8 + c9 + c10
	fmt.Println(count)
}
