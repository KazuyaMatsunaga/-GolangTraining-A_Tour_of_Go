package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)

	for i := range words {
		if m[words[i]] == 0 {
			for j := range words {
				if words[i] == words[j] {
					m[words[i]] += 1
				}
			}
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
