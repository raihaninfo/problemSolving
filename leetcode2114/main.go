package main

import (
	"fmt"
	"strings"
)

func main() {
	s := mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})
	fmt.Println(s)
}

func mostWordsFound(sentences []string) int {
	maxlen := 0
	for _, s := range sentences {
		slen := strings.Split(s, " ")
		if len(slen) > maxlen {
			maxlen = len(slen)
		}
	}
	return maxlen
}
