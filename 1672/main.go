package main

import "fmt"

func main() {
	acc := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(acc))
}

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, row := range accounts {
		wealth := 0
		for _, v := range row {
			wealth += v
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
