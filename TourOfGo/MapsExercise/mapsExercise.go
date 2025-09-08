package main

import (
	"fmt"
	"strings"
)

// Implement WordCount. It should return a
// map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against
// the provided function and prints success or failure.

// You might find strings.Fields helpful.

var sliceOfWords []string = []string{
	"Hello",
	"Jason",
	"It",
	"should",
	"return",
	"a",
	"map",
	"of",
	"the",
	"counts",
	"each",
	"word",
}

func WordCount(s string) map[string]int {
	sliceOfStrs := strings.Split(s, " ")
	m := make(map[string]int)

	for _, str := range sliceOfStrs {
		m[str] = len(str)
	}

	return m
}

func main() {
	s := strings.Join(sliceOfWords, " ")
	fmt.Println(s)
	fmt.Println(WordCount(s))
	fmt.Println(len(WordCount(s)))
}
