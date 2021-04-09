package main

import "fmt"

func main() {
	arr := []int{4, 7, -10, 6, -12, 3}
	// arr := []int{10, -20, 15, 30, -40}
	fmt.Println(nonNegativeSumOfArray(arr))
}

func nonNegativeSumOfArray(arr []int) int {
	minRes, sum := 0, 0
	for _, v := range arr {
		sum += v
		if sum < 0 && sum < minRes {
			minRes = sum
		}
	}
	return -(minRes - 1)
}
