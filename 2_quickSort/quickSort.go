package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func choosePivot(input []int, left, right int) int {
	// use the first element as the pivot
	// return left

	// use the last element as the pivot
	// return right

	// use the median of three as the pivot
	return medianOfThree(input, left, right)
}

func medianOfThree(input []int, left, right int) int {
	mid := left + ((right - left) / 2)
	if input[left] < input[mid] {
		if input[left] > input[right] {
			return left
		}
		if input[right] < input[mid] {
			return right
		}
		return mid
	}
	if input[mid] > input[right] {
		return mid
	}
	if input[left] < input[right] {
		return left
	}
	return right
}

func partition(input []int, left, right int) int {
	p := input[left]
	i := left + 1
	for j := left + 1; j <= right; j++ {
		if input[j] < p {
			temp := input[j]
			input[j] = input[i]
			input[i] = temp
			i++
		}
	}
	temp := input[left]
	input[left] = input[i-1]
	input[i-1] = temp
	return i - 1
}

// QuickSort sorts an array and returns the number of comparisons carried out.
func QuickSort(input []int, left, right int) int {
	if left > right {
		return 0
	}
	i := choosePivot(input, left, right)

	// swap pivot to leftmost position
	temp := input[left]
	input[left] = input[i]
	input[i] = temp

	// new pivot position
	j := partition(input, left, right)

	comparisons := right - left
	leftComps := QuickSort(input, left, j-1)
	rightComps := QuickSort(input, j+1, right)

	return comparisons + leftComps + rightComps
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("file argument missing")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var nums [10008]int
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var i int
	for scanner.Scan() {
		nums[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	numSlice := nums[:i]
	fmt.Printf("number of comparisons: %d\n", QuickSort(numSlice, 0, i-1))
}
