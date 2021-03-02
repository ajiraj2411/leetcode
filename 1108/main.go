package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	ip := "1.1.1.1"
	fmt.Println(defangIPaddr(ip))
	fmt.Println(defangIPaddrWithOutInBuidLibrary(ip))
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func defangIPaddrWithOutInBuidLibrary(address string) string {
	var s bytes.Buffer
	for _, c := range address {
		if c == '.' {
			s.WriteString("[.]")
		} else {
			s.WriteRune(c)
		}
	}
	return s.String()
}
