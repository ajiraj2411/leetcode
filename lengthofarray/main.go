package main

import "fmt"

func main() {
	a := []int{1, 2}
	fmt.Println(solve(a))
}

func solve(nums []int) bool {
	// Write your code here
	if len(nums) == 1 {
		return true
	}
	return false
}
