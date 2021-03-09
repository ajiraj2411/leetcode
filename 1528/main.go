package main

import "fmt"

func main() {
	arr := []int{4, 0, 1, 3, 2}
	fmt.Println(restoreString("ajith", arr))
}

func restoreString(s string, indices []int) string {
	ans := make([]byte, len(s))
	for index, ch := range []byte(s) {
		newIndex := indices[index]
		ans[newIndex] = ch
	}
	return string(ans)
}
