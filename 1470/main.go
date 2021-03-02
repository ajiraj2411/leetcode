package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(shuffle(nums, 2))
}
func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i, j := 0, 0; i < n; i, j = i+1, j+2 {
		res[j], res[j+1] = nums[i], nums[i+n]
	}
	return res
}
