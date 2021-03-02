package main

import "fmt"

func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	fmt.Println(countMatches(items, ruleKey, ruleValue))
	fmt.Println(countMatchesFaster(items, ruleKey, ruleValue))
}

var fields = map[string]int{
	"type":  0,
	"color": 1,
	"name":  2,
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	for _, item := range items {
		if ruleValue == item[fields[ruleKey]] {
			count++
		}
	}
	return count
}

func countMatchesFaster(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	if ruleKey == "type" {
		for _, item := range items {
			if ruleValue == item[0] {
				count++
			}
		}
	} else if ruleKey == "color" {
		for _, item := range items {
			if ruleValue == item[1] {
				count++
			}
		}
	} else {
		for _, item := range items {
			if ruleValue == item[2] {
				count++
			}
		}
	}
	return count
}
