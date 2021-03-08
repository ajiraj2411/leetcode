package main

import "fmt"

func main() {
	arr := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(arr))
}

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		count := 0
		for j := range nums {
			if nums[j] < nums[i] {
				count++
			}
		}
		res[i] = count
	}
	return res
}
