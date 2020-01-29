package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mergeAndCountSplitInv(arrA, arrB []int) ([]int, int) {
	var i, j, splitInv int
	n := len(arrA) + len(arrB)
	output := make([]int, n)
	for k := 0; k < n; k++ {
		if i < len(arrA) && j < len(arrB) {
			if arrA[i] <= arrB[j] {
				output[k] = arrA[i]
				i++
			} else {
				output[k] = arrB[j]
				j++
				splitInv += len(arrA) - i
			}
		} else {
			if i < len(arrA) {
				output[k] = arrA[i]
				i++
			} else {
				output[k] = arrB[j]
				j++
			}
		}
	}
	return output, splitInv
}

// MergeSort sorts an int array and counts the inversions.
func MergeSort(input []int) ([]int, int) {
	if len(input) < 2 {
		return input, 0
	}
	arrA, lInv := MergeSort(input[:len(input)/2])
	arrB, rInv := MergeSort(input[len(input)/2:])
	arr, splitInv := mergeAndCountSplitInv(arrA, arrB)
	return arr, (lInv + rInv + splitInv)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("file argument missing")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var nums [100008]int
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
	_, inversions := MergeSort(nums[:i])
	fmt.Printf("number of inversions: %d\n", inversions)
}
