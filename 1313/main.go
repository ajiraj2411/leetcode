package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(input))
}

func decompressRLElist(nums []int) []int {
	// result := []int{}
	length := 0
	for i := 0; i < len(nums); i += 2 {
		length += nums[i]
	}
	result := make([]int, length)
	idx := 0
	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			result[idx] = val
			idx++
		}
	}
	return result
}
