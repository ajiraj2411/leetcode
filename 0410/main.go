package main

import "fmt"

func main() {
	array := []int{4, 1, 2, 1, 2}
	numOfSplit := 2
	fmt.Println(largestSum(array, numOfSplit))
}

func largestSum(nums []int, m int) int {
	eqSplit := len(nums) / m
	a := sumOfArray(nums[:eqSplit])
	return a
}

func sumOfArray(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
