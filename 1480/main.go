package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(runningSum(arr))
}

func runningSum(nums []int) []int {
	for i, num := range nums[1:] {
		fmt.Println(i, num)
		nums[i+1] = nums[i] + num
	}
	return nums
}
