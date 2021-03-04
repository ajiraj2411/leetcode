package main

import (
	"fmt"
)

func main() {
	// arr := []int{8, 7, 25, 4, 12, 40, 97}
	// arr := []int{3, 1, 5, 8}
	// arr := []int{3, 1, 5}
	arr := []int{9, 76, 64, 21, 97, 60}
	// arr := []int{1, 5}
	fmt.Println(maxCoins(arr))
}

func maxCoins(A []int) int {
	A = append(append([]int{1}, A...), 1)
	mem := map[[2]int]int{}
	var dfs func(_, _ int) int
	dfs = func(i, j int) int {
		if ret, ok := mem[[2]int{i, j}]; ok {
			return ret
		}
		out := 0
		for k := i + 1; k <= j-1; k++ {
			ret := A[i]*A[k]*A[j] + dfs(i, k) + dfs(k, j)
			if ret > out {
				out = ret
			}
		}
		mem[[2]int{i, j}] = out
		return out
	}
	return dfs(0, len(A)-1)
}
