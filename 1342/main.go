package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(123))
}

func numberOfSteps(num int) int {
	res := 0
	for num != 0 {
		if num%2 == 0 {
			num >>= 1
		} else {
			num--
		}
		res++
	}
	return res
}
