// @see: https://leetcode.com/problems/implement-trie-prefix-tree
package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, c := range word {
		_, exists := node.children[c]
		if !exists {
			node.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, c := range word {
		_, exists := node.children[c]
		if !exists {
			return false
		}
		node = node.children[c]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, c := range prefix {
		_, exists := node.children[c]
		if !exists {
			return false
		}
		node = node.children[c]
	}
	return true
}

func main() {
	trie := NewTrie()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
}
