package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	numMaps := map[int]int{}
	for _, n := range nums {
		numMaps[n] = n
	}
	return len(nums) != len(numMaps)
}
