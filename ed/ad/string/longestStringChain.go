// Longest String Chain.
//
// A Magic Word Chain (MWC) is a sequence of words where each step in the
// sequence inserts a single character into the preceding word.
//
// Given a list of words which can be used, we want to find and return the
// longest possible Magic Word Chain (MWC).
//
// Example: From the list of words:
//   [I, IN, SIN, SING, SIGN, STING, SIGNS, STRING, STRINGS]
// We can make the chains:
//   I -> IN -> SIN -> SIGN -> SIGNS (length 5)
//   I -> IN -> SIN -> SING -> STING -> STRING -> STRINGS (length 7)

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"I", "IN", "SIN", "SING", "SIGN", "STING", "SIGNS", "STRING", "STRINGS"}
	r1 := getLongestChainLength(input)
	r2 := getLongestChain(input)
	fmt.Printf("\n %v | %v \n", r1, r2)
}

func getLongestChain(words []string) []string {
	// Sort words by word length.
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	// Create chains of words.
	chains := make(map[string][]string)
	for _, word := range words {
		chains[word] = []string{word}
		for _, prevWord := range words {
			if len(word) == len(prevWord)+1 && strings.ContainsAny(word, prevWord) {
				wordChain := chains[word]
				prevWordChain := chains[prevWord]
				if len(prevWordChain)+1 > len(wordChain) {
					newChain := make([]string, len(prevWordChain))
					copy(newChain, prevWordChain)
					newChain = append(newChain, word)
					chains[word] = newChain
				}
			}
		}
	}

	// Get longest value.
	res := []string{}
	for _, v := range chains {
		if len(v) > len(res) {
			res = v
		}
	}

	return res
}

func getLongestChainLength(words []string) int {
	// Sort words by word length.
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	// Calculate lengths for possible chains.
	maxLengths := make(map[string]int)
	for _, word := range words {
		maxLengths[word] = 1
		for _, prevWord := range words {
			if len(word) == len(prevWord)+1 && strings.ContainsAny(word, prevWord) {
				maxLengths[word] = max(maxLengths[word], maxLengths[prevWord]+1)
			}
		}
	}

	// Get max value.
	res := 0
	for _, v := range maxLengths {
		if v > res {
			res = v
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
