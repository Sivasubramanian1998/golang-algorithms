package main

import "fmt"

//Alphabet is the number of possible characters in the Trie
const Alphabet = 26

//Node represent each node in a Trie
type Node struct {
	children [26]*Node
	isEnd    bool
}

//Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

//InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//Search will take a word and return true if that word is included in a trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := InitTrie()
	myTrie.Insert("siva")
	fmt.Println(myTrie.Search("siva"))
}
