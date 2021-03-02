package main

import "fmt"

func main() {
	n := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(n))
}

func singleNumber(nums []int) int {
	result := 0
	for num := range nums {
		result ^= num
	}
	return result
}
