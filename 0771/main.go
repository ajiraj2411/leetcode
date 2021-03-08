package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(jewels string, stones string) int {
	res := 0
	for _, stone := range stones {
		for _, jewel := range jewels {
			if stone == jewel {
				res++
			}
		}
	}
	return res
}
