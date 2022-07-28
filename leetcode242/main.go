package main

import (
	"fmt"
)

func main() {
	jj := isAnagram("cat", "tat")
	fmt.Println(jj)
}

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true

}
