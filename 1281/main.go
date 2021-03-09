package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subtractProductAndSum(143))
}

func subtractProductAndSum(n int) int {
	prod, sum := 1, 0
	for _, v := range strconv.Itoa(n) {
		i, _ := strconv.Atoi(string(v))
		prod *= i
		sum += i
	}
	return prod - sum
}
