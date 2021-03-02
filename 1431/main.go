package main

import "fmt"

func main() {
	arr := []int{12, 1, 12}
	fmt.Println(kidsWithCandies(arr, 10))
}
func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := []bool{}
	maxVal := -1
	for _, value := range candies {
		if value > maxVal {
			maxVal = value
		}
	}
	for _, candy := range candies {
		res = append(res, candy+extraCandies >= maxVal)
	}

	return res
}
