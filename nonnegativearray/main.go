package main

import "fmt"

func main() {
	arr := []int{4, 7, -10, 6, -12, 3}
	fmt.Println(nonNegativeSumOfArray(arr))
}

func nonNegativeSumOfArray(arr []int) int {
	minRes, sum := 0, 0
	for _, v := range arr {
		sum += v
		if sum < 0 {
			minRes -= sum
		}
	}
	return minRes
}
