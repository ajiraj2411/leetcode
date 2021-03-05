package main

import "fmt"

func main() {
	// arr := []int{1, 2, 3, 1, 1, 3}
	arr := []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(arr))

}

func numIdenticalPairs(nums []int) int {
	pair := 0
	hMap := make(map[int]int)
	for _, num := range nums {
		pair += hMap[num]
		hMap[num]++
	}
	return pair
}
