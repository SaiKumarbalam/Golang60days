package main

import (
	"fmt"
	"sort"
)

func main() {
	// create an unsorted slice of integers
	nums := []int{4, 2, 8, 5, 1, 9}

	// use the sort package to sort the slice
	sort.Ints(nums)

	// print the sorted slice
	fmt.Println(nums)
}
