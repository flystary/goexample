package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindromeString(s string) bool {
	compile := regexp.MustCompile(`[\w]+`)
	findstring := compile.FindAllString(s, -1)
	a := strings.ToUpper(strings.TrimSpace(strings.Join(findstring, "")))
	return isPalind(a)
}

func isPalind(s string) bool {
	l, f := 0, len(s)-1
	for l < f {
		if s[1] == s[f] {
			l++
			f--
			continue
		}
		return false
	}
	return true
}

func main() {
	fmt.Println(isPalindromeString("race a car"))
}
