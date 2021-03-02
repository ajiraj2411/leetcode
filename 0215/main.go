package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{98, 4, 76, 2, 65, 10, 23, 45}
	fmt.Println(findKthLargest(arr, 3))
}
func findKthLargest(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]
}
