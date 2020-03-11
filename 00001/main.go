package main

import "fmt"

func main() {
	input1 := []int{2, 3, 5, 7, 11, 13}
	input2 := 12
	fmt.Println(twoSum(input1, input2))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		compliment := target - num
		if v, ok := m[compliment]; ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return nil
}
